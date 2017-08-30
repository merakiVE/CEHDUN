# BMKero's

El estandar de programacón definido por el equipo de desarrollo tiene diferentes reglas para facilitar el desarrollo grupal y la compatibilidad entre el trabajo realizado por cada miembro del equipo.
Las siguientes reglas son las tomadas en cuenta como estandar de realización de este sitio web.

### Funcionalidad de las carpetas.

**_public_**

    Esta carpeta contiene los archivos que el servidor va a servir. Para este proyecto contendra nuestro index.html y el favicon.ico.

**_src_**

    Tan solo contendra la carpeta components y los archivos de arranque de nuestra aplicacion como lo puede ser nuestro index.html.

**_components_**

    Aqui se ubicaran todos los componentes que hayamos creado. Es importante tener en cuenta varias reglas para crear un nuevo componente:

    1. Los nombres de los componententes deben ser representativos segun su funcionalidad.

    2. El nombre del componente debe iniciar con la primera letra en mayuscula, ejemplo “Component”.

    3. Cada componente debe estar contenido en una carpeta cuyo nombre debe ser identico al de el componente que contendra.

    4. En la carpeta que contendra nuestro componente se deberan ubicar los archivos que usara nuestro componente para complementar su funcionalidad , talez como las hojas de estilo en cascada( css ) y las imagenes.

    5. El nombre de los archivos de estilo ( css ) debe ser tal cual es el del componente al que va asociado.

### Aplicación de los estilos.

    Para este proyecto definimos aplicar estilos a los elementos definidos mediante la importanción de un archivo de hojas de estilo en casacada de la siguiente manera:

    ```javascript
        import './style.css';
    ```

    Los puntos a tener en cuenta para aplicar estilos en este proyecto son los siguientes:

    1. Se deben aplicar estilos mediante el uso de clases y id.

    2. No se debe aplicar estilos mediante el uso de las etiquetas.

    3. La organización de los estilos en nuestra hoja de estilo en cascada debe ser la siguiente, en primer lugar debemos definir las ```media-query```, seguido de las  ```class```, pseudo clases y por ultimo los ```id```.

    4. Para aplicar clases o ids se debe hacer de la siguiente manera, los nombres de las clases o los id deben iniciar con el nombre del componente al que estaran relacionados seguido de un ```–``` y el nombre de la funcionalidad que tendra dicho elemento.

    **Ejemplo**

    ```css
    .nameComponent-funcionalidad {
        color: red;
    }
    #nameComponent-funcionalidad {
        background-color: darkred;
    }
    ```