package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/semconv/v1.17.0"  
)

func main() {

    cleanup := initTracer()
    defer cleanup()

    router := gin.Default()
    router.POST("/order", handleOrder)

    log.Println("Servidor rodando em http://localhost:8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("falha ao iniciar o servidor: %v", err)
    }
}

// initTracer configura o Jaeger com atributos de recurso que servem como "process" para descrever o ambiente de execução
func initTracer() func() {
    exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
    if err != nil {
        log.Fatalf("falha ao criar exportador Jaeger: %v", err)
    }

    provider := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            attribute.String("service.name", "order-service"),
            attribute.String("service.environment", "development"),
            attribute.String("service.version", "1.0.0"),
        )),
    )
    otel.SetTracerProvider(provider)
    return func() { _ = provider.Shutdown(context.Background()) }
}

// handleOrder é o endpoint que processa o pedido
func handleOrder(c *gin.Context) {
    
    var request struct {
        ProductID string  `json:"product_id"`
        Amount    float64 `json:"amount"`
        CardType  string  `json:"card_type"`
        Quantity  int     `json:"quantity"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de entrada inválidos"})
        return
    }

    ctx := context.Background()

    withdrawFromInventory(ctx, request.ProductID, request.Quantity)
    processPayment(ctx, request.ProductID, request.CardType, request.Amount)
    dispatchOrder(ctx, request.ProductID)

    c.JSON(http.StatusOK, gin.H{"status": "Pedido processado com sucesso!"})
}

// withdrawFromInventory simula a retirada de um item do estoque
func withdrawFromInventory(ctx context.Context, productID string, quantity int) {
    tracer := otel.Tracer("order-service")
    _, span := tracer.Start(ctx, "withdrawFromInventory")
    defer span.End()

    // Estoque
    var inventory = 0
    var initialStock = 100
    span.SetAttributes(attribute.Int("inventory.initialStock", initialStock))

    // Simula a retirada do estoque
    if quantity <= initialStock {
        inventory = initialStock - quantity
        fmt.Printf("Retirando %d unidades do produto %s do estoque\n", quantity, productID)
    } else {
        fmt.Printf("Estoque insuficiente para o produto %s\n", productID)
    }

    span.SetAttributes(attribute.Int("inventory.orderQtd", quantity))

    var updatedStock = inventory
    span.SetAttributes(attribute.Int("inventory.updatedStock", updatedStock))
}

// processPayment simula o processamento do pagamento com cartão
func processPayment(ctx context.Context, productID, cardType string, amount float64) {
    tracer := otel.Tracer("order-service")
    _, span := tracer.Start(ctx, "processPayment")
    span.SetAttributes(
        attribute.String("product.id", productID),
        attribute.String("payment.cardType", cardType),
        attribute.Float64("payment.amount", amount),
    )
    defer span.End()

    fmt.Printf("Processando pagamento de %.2f para o produto %s com o cartão %s\n", amount, productID, cardType)
}

// dispatchOrder simula o despacho do pedido
func dispatchOrder(ctx context.Context, productID string) {
    tracer := otel.Tracer("order-service")
    _, span := tracer.Start(ctx, "dispatchOrder")
    span.SetAttributes(attribute.String("product.id", productID))
    defer span.End()

    fmt.Printf("Despachando o pedido para o produto %s\n", productID)
}

