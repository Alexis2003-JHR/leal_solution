# DIAGRAMAS DE FLUJO POR CASO DE USO

## Configurar campañas para un comercio y sucursales
```mermaid
sequenceDiagram
    participant Usuario
    participant API
    participant BaseDatos
    participant Comercio

    Usuario->>API: Enviar datos de la campaña
    API->>BaseDatos: Validar y guardar campaña
    BaseDatos-->>API: Confirmación de guardado
    API->>Comercio: Notificar nueva campaña
    Comercio-->>API: Confirmación de recepción
    API-->>Usuario: Respuesta de éxito o error
```
## Consultar campañas de un comercio y sucursal
```mermaid
sequenceDiagram
    participant Usuario
    participant API
    participant BaseDatos

    Usuario->>API: Solicitar campañas de un comercio y sucursal
    API->>BaseDatos: Consultar campañas activas
    BaseDatos-->>API: Retornar campañas encontradas (o vacío)
    API-->>Usuario: Enviar respuesta con campañas o mensaje vacío
```

## Realizar acumulación de puntos o cashback

```mermaid
sequenceDiagram
    participant Usuario
    participant API
    participant BaseDatos
    participant Notificaciones

    Usuario->>API: Realizar compra
    API->>BaseDatos: Validar compra y usuario
    API->>BaseDatos: Verificar si hay campaña activa
    BaseDatos-->>API: Respuesta con reglas de campaña (si aplica)
    API->>BaseDatos: Calcular y almacenar puntos/cashback
    BaseDatos-->>API: Confirmación de acumulación
    API->>Notificaciones: Enviar notificación de acumulación al usuario
    Notificaciones-->>Usuario: Confirmación de puntos/cashback
```

## Redimir puntos o cashback en un comercio

```mermaid
sequenceDiagram
    participant Usuario
    participant API
    participant BaseDatos
    participant Comercio

    Usuario->>API: Solicitar redención de puntos/cashback
    API->>BaseDatos: Verificar saldo del usuario
    BaseDatos-->>API: Confirmación de saldo suficiente / error
    API->>BaseDatos: Aplicar redención y actualizar saldo
    BaseDatos-->>API: Confirmación de redención
    API->>Comercio: Notificar transacción de redención
    Comercio-->>API: Confirmación de recepción
    API-->>Usuario: Confirmación de redención o error
```