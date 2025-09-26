Taller Práctico: Creación de Paquetes en Go

Este repositorio contiene el desarrollo del *Taller Práctico de Creación de Paquetes*,  
donde se implementan dos funcionalidades principales en Go:

---

 Funcionalidades
 1. Conversor de monedas
El usuario ingresa un valor en dólares (USD) y selecciona la moneda a la que desea convertir.  

Las monedas permitidas son:
- Euros (EUR)  
- Libras Esterlinas (LB)  
- Won Surcoreano (WON)  
- Bitcoin (BTC)  

Ejemplo de uso:
```go
resultado := utilidades.Conversor(100, 1) // Convierte 100 USD a Euros
fmt.Println("100 USD en Euros:", resultado)
2. Contador de Vocales
El usuario ingresa una frase y el programa devuelve el número de veces que aparece cada vocal.

Ejemplo de uso:

go
Copy code
frase := "Hola Mundo"
conteo := utilidades.ContadorVocales(frase)
fmt.Println(conteo) // a:1, e:0, i:0, o:2, u:1
 Estructura del Proyecto
bash
Copy code
paquete_utilidades/
│── utilidades.go   # Funciones del conversor y el contador
│── go.mod          # Configuración del módulo
│── README.md       # Documentación del taller
│── .gitignore      # Archivos ignorados por Git
 Uso desde un nuevo proyecto
Crear un nuevo proyecto en Go:

bash
Copy code
go mod init ejemplo.com/mi-proyecto
Importar el paquete desde GitHub:

go
Copy code
import "github.com/Mimi55-gif/paquete-utilidades-emiliaarmas"
Ejecutar el programa y probar las funciones.

 Comandos utilizados
bash
Copy code
# Inicializar repositorio
git init

# Agregar archivos
git add .

# Crear commit
git commit -m "Primer commit: paquete utilidades"

# Vincular con GitHub
git remote add origin https://github.com/Mimi55-gif/paquete-utilidades-emiliaarmas.git

# Subir a GitHub
git push -u origin main

Autor Emilia Armas
