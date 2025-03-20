# Análisis del Problema

## 1. Introducción
Leal es una plataforma de fidelización que permite a los usuarios ganar **puntos** o **cashback** por sus compras en comercios afiliados. Los comercios pueden configurar **factores de conversión** y establecer **campañas promocionales** que alteren la acumulación de recompensas.

### **Objetivo principal:**
- Permitir que los usuarios acumulen y rediman puntos o cashback en los comercios de la red.  
- Brindar flexibilidad a los comercios para definir reglas de fidelización personalizadas.  
- Permitir campañas temporales que alteren la forma en que los usuarios acumulan recompensas.  

## 2. Reglas de Negocio Identificadas

### **Reglas de acumulación de recompensas**
1. **Cada comercio define su propio factor de conversión**
2. **Los puntos solo pueden ser usados en el comercio donde fueron obtenidos**
3. **El cashback puede ser usado en cualquier comercio de la red Leal**
4. **Las compras registradas deben aplicar las reglas de acumulación del comercio correspondiente**

### **Reglas de campañas promocionales**
5. **Las campañas tienen una vigencia definida (fecha de inicio y fin)**
6. **Las campañas pueden ofrecer multiplicadores o beneficios adicionales**
7. **Las campañas son específicas para cada sucursal**

### **Reglas de redención de puntos y cashback**
8. **Cada comercio define una tabla de premios para la redención de puntos**
9. **El cashback se redime de forma directa, equivalente a dinero (1 Coin = $1)**
10. **Debe haber suficiente saldo de puntos/cashback para redimir**

## 3. Consideraciones Técnicas

### **Persistencia de Datos**
- Se debe mantener un histórico de transacciones.
- Las campañas deben registrarse en la BD.
- El saldo de puntos y cashback debe calcularse dinámicamente o pre-calcularse.

### **Escalabilidad y Performance**
- Consultas eficientes para calcular saldo.
- Uso de caché para mejorar tiempos de respuesta.
- Manejo de concurrencia en redenciones.

### **Seguridad y Prevención de Fraude**
- Registro de auditoría.
- Autenticación y autorización.
- Evitar manipulación de saldo.

## 4. Casos de Uso Clave

### **1. Acumulación de recompensas**
1. Un usuario realiza una compra en una sucursal.
2. Se verifica si hay una campaña activa.
3. Se calcula la cantidad de puntos o cashback.
4. Se registra la transacción y se actualiza el saldo del usuario.

### **2. Redención de recompensas**
1. Un usuario quiere redimir puntos o cashback en un comercio.
2. Se valida si tiene suficiente saldo.
3. Se descuenta la cantidad correspondiente.
4. Se registra la redención en el sistema.

### **3. Configuración de campañas**
1. Un comercio define una campaña en una sucursal.
2. Se establecen fechas, montos mínimos y multiplicadores.
3. Se guarda en la base de datos.

### **4. Consulta de saldo de puntos o cashback**
1. Un usuario consulta cuánto ha acumulado.
2. El sistema calcula el total basado en las transacciones registradas.

## 5. Endpoints del API

### **Usuarios**
- `POST /users` → Crear un usuario
- `GET /users/{id}/balance` → Consultar saldo de puntos y cashback

### **Compras y Recompensas**
- `POST /transactions` → Registrar una compra y calcular recompensas
- `GET /transactions/{id}` → Consultar detalles de una transacción

### **Redención de Recompensas**
- `POST /redemptions` → Redimir puntos o cashback

### **Campañas**
- `POST /campaigns` → Crear una nueva campaña
- `GET /campaigns/{commerce_id}` → Consultar campañas activas para un comercio

### **Comercios y Sucursales**
- `POST /merchants` → Crear un nuevo comercio
- `POST /branches` → Crear una sucursal para un comercio

## 6. Conclusión
 Aspectos más importantes a considerar:
- **Definir bien la lógica de acumulación**.
- **Diferenciar puntos y cashback claramente**.
- **Manejo seguro de saldo**.
