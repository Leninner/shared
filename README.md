# Shared Module

Módulo compartido para TanEats que contiene tipos, utilidades y componentes comunes utilizados por todos los microservicios.

## Estructura del Proyecto

```
shared/
├── domain/
│   ├── entity/
│   │   ├── aggregate_root.go
│   │   └── base_entity.go
│   ├── event/
│   │   ├── domain_events.go
│   │   └── publisher/
│   │       └── domain_event_publisher.go
│   ├── exception/
│   │   └── domain_exception.go
│   └── valueobject/
│       ├── base_id.go
│       ├── customer_id.go
│       ├── money.go
│       ├── order_approval_status.go
│       ├── order_id.go
│       ├── order_status.go
│       ├── payment_status.go
│       ├── product_id.go
│       └── restaurant_id.go
├── logger/
│   ├── factory.go
│   └── logger.go
├── utils/
│   └── validator/
│       └── validator.go
├── go.mod
└── README.md
```

## Tecnologías

- **Go 1.24.2**
- **Uber Zap** - Logging estructurado
- **Google UUID** - Generación de IDs únicos

## Desarrollo con GitHub Submódulos

### Configuración Inicial

Este módulo es parte de un proyecto modular usando Git submódulos. El repositorio principal está en: `https://github.com/leninner/taneats`

### Clonar el Proyecto Completo

```bash
git clone --recursive https://github.com/leninner/taneats.git
cd taneats/shared
```

### Trabajar en el Módulo Compartido

1. **Navegar al directorio del módulo:**
   ```bash
   cd shared
   ```

2. **Hacer cambios en el código**

3. **Commit y push de cambios:**
   ```bash
   git add .
   git commit -m "feat: add new value object for order tracking"
   git push origin main
   ```

4. **Actualizar el submódulo en el repositorio principal:**
   ```bash
   cd ..  # volver al root del proyecto
   git add shared
   git commit -m "Update shared submodule"
   git push origin main
   ```

### Actualizar Dependencias

```bash
go mod tidy
go mod download
```

### Ejecutar Tests

```bash
go test ./...
```

## Componentes Compartidos

### Value Objects
- `BaseID` - ID base para todas las entidades
- `CustomerID` - Identificador único de cliente
- `OrderID` - Identificador único de orden
- `RestaurantID` - Identificador único de restaurante
- `ProductID` - Identificador único de producto
- `Money` - Tipo para manejo de dinero
- `OrderStatus` - Estados de una orden
- `PaymentStatus` - Estados de pago
- `OrderApprovalStatus` - Estados de aprobación

### Entidades Base
- `BaseEntity` - Entidad base con ID y timestamps
- `AggregateRoot` - Raíz de agregado para DDD

### Eventos de Dominio
- `DomainEvent` - Evento base de dominio
- `DomainEventPublisher` - Publicador de eventos

### Utilidades
- `Logger` - Logger estructurado con Zap
- `Validator` - Utilidades de validación

## Uso en Otros Servicios

### Importar el módulo

```go
import "github.com/leninner/shared"
```

### Ejemplo de uso

```go
package main

import (
    "github.com/leninner/shared/domain/valueobject"
    "github.com/leninner/shared/logger"
)

func main() {
    // Crear un ID de orden
    orderID := valueobject.NewOrderID()
    
    // Usar el logger
    log := logger.NewLogger()
    log.Info("Order created", "orderID", orderID.String())
}
```

## Convenciones de Git

### Commits
- `feat:` nuevas características
- `fix:` correcciones de bugs
- `refactor:` refactorización de código
- `test:` agregar o modificar tests
- `docs:` documentación
- `breaking:` cambios que rompen compatibilidad

### Branches
- `main` - código estable
- `develop` - desarrollo activo
- `feature/` - nuevas características
- `hotfix/` - correcciones urgentes

## Versionado

Este módulo sigue Semantic Versioning (SemVer):
- `MAJOR.MINOR.PATCH`
- Cambios breaking incrementan MAJOR
- Nuevas características incrementan MINOR
- Bug fixes incrementan PATCH

## Impacto de Cambios

⚠️ **Importante**: Los cambios en este módulo afectan a todos los servicios que lo utilizan.

### Antes de hacer cambios:
1. Verificar qué servicios dependen de este módulo
2. Considerar compatibilidad hacia atrás
3. Comunicar cambios breaking a otros equipos
4. Actualizar documentación

### Después de cambios:
1. Actualizar versiones en todos los servicios
2. Ejecutar tests en todos los servicios
3. Verificar que no hay regresiones

## Troubleshooting

### Problema: Submódulo no actualizado
```bash
git submodule update --remote shared
```

### Problema: Dependencias no encontradas
```bash
go mod tidy
go mod download
```

### Problema: Tests fallando
```bash
go clean -testcache
go test ./...
```

### Problema: Conflictos de versiones
```bash
go mod graph | grep shared
``` 