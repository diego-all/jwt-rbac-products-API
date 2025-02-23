# Login


Práctica	Descripción

✅ No exponer datos sensibles	Devuelve solo el token, sin información del usuario.

✅ Usar Authorization: Bearer	Nunca enviar el token en URL o cuerpo.

✅ Expiración corta para JWT	Máximo 15-30 min, no tokens de larga duración.

✅ Refresh Tokens para sesiones largas	Solo se usa para renovar JWTs.

✅ Firmar JWTs con HS256 o RS256	Nunca exponer la clave secreta.

✅ Verificar tokens en cada request	No confiar en tokens sin validación.

✅ Lista negra de tokens revocados	Para invalidar tokens antes de su expiración.



