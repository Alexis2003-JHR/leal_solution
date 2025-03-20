# Contexto

### Leal

Leal es una startup que ayuda a las personas a obtener recompensas a través de sus compras del día a día y que, a su vez, genera ventas incrementales a los comercios que adoptan el modelo.  

Como líneas de fidelización, ofrecemos a los usuarios la posibilidad de adquirir **Puntos Leal** y/o **Cashback** como retribución a sus compras en nuestras marcas y comercios, teniendo en cuenta las siguientes condiciones:  

 ### Puntos Leal

- El total de puntos a ganar se calcula bajo el factor de conversión que define el comercio sobre el valor de la compra.  
  - Un punto es equivalente a un valor en pesos.  
  - Ejemplo: la marca Texaco define que otorgará al usuario **1 punto por cada $1.000** tanqueados en gasolina.  
  - Una compra de **$20.000** le entregará al usuario **20 puntos**.  
    - Solamente pueden ser usados en el comercio donde fueron ganados.  
    - La redención de los puntos se da bajo la tabla de premios que el comercio defina.  
    - Ejemplo: Texaco define que para ganar una carga de gasolina por el valor de **$5.000**, el usuario debe contar con **100 puntos**.  

### Cashback

- El total de **Coins** a ganar se calcula bajo las mismas condiciones que los puntos.  
- Pueden ser usados en cualquier comercio de la red Leal.  
- La equivalencia en términos de redención es **1 Leal Coin = $1**.  


# Premisa

Con motivo de la apertura de dos nuevas estaciones de servicio, Texaco quiere
premiar a los usuarios que hagan compras en estas sucursales. Para esto es
necesario crear una campaña especial que le permita al usuario acumular puntos o
cashback adicionales de la siguiente manera:
- **Sucursal 1:** por todas las compras realizadas entre el 15 y 30 de mayo los
usuarios acumularan el doble de puntos o cashback.
- **Sucursal 2:** por todas las compras realizadas superiores a $20.000 entre el
15 y 20 de mayo los usuarios acumularán un 30% adicional de puntos o
cashback.

# Requerimientos

- Configurar campañas para un comercio y sucursales.
- Consultar campañas de un comercio y sucursal.
- Realizar acumulación de puntos o cashback si el comercio y la sucursal tiene
configurada una campaña.
- Redimir los puntos o cashback en un comercio. **(opcional)**

    ## Requerimientos técnicos

- Se debe usar una base de datos de su elección
- Se debe enviar un archivo readme en el cual se especifique como levantar
el proyecto
- Utilizar arquitectura hexagonal
- Construir pruebas unitarias
- Se debe manejar el estándar http
- Construir swagger con la documentación de los endpoints
    ### Deseable
- Implementar migraciones de base de datos
