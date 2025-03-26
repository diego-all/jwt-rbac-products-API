func (a *App) login(w http.ResponseWriter, r *http.Request) {
req := database.Login{}

if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
a.log.Errorf("Error decode request: %s\n", err)
w.WriteHeader(http.StatusBadRequest)
return
}

user, err := database.FetchUserByLogin(a.Db, req.Login)
if err != nil {
a.log.Errorf("Error get user from database: %s\n", err)
w.WriteHeader(http.StatusBadRequest)
return
}

if user.IsBlocked {
a.log.Errorf("The user %d is blocked!\n", user.ID)
w.WriteHeader(http.StatusForbidden)
return
}

if lockout.IsLockedOut(a.Db, user.ID) {
a.log.Errorf("The user %d is locked out!\n", user.ID)
w.WriteHeader(http.StatusBadRequest)
return
}

match := auth.VerifyPassword(user.Hash, req.Password)

if !match {
err := database.RegisterFailedAttempt(a.Db, user.ID, r.RemoteAddr)
if err != nil {
a.log.Errorf("Wrong password for user '%s' from IP: %s\n",
user.Login, r.RemoteAddr)
}

err = lockout.RegisterFailedAttempt(a.Db, user.ID)
if err != nil {
a.log.Errorf("Error register failed attempt %s\n", err)
a.InternalServerErrorWriter(w)
return
}

w.WriteHeader(http.StatusUnauthorized)
return
}

err = lockout.RegisterGoodAttempt(a.Db, user.ID)
if err != nil {
a.log.Errorf("Error register good attempt %s\n", err)
a.InternalServerErrorWriter(w)
return
}

token := uuid.NewV4()

err = a.redis.Put(&redis.Item{
Key:   token.String(),
Value: []byte(strconv.Itoa(int(user.ID))),
Exp:   TOKEN_AUTH_LIFETIME,
})

if err != nil {
a.log.Errorf("Error put token to redis: %s\n", err)
a.InternalServerErrorWriter(w)
return
}

const prefix = "Bearer "
w.Header().Set("Authorization", prefix+token.String())
}