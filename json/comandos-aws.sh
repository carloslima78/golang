
# Executar o container LocalStack:
docker run --rm -d -p 4566:4566 -p 4571:4571 --name localstack localstack/localstack

# Listar as imagens
docker image ls

# Criar o bucket S3
aws --endpoint-url=http://localhost:4566 s3 mb s3://meu-bucket

# Listar os buckets S3
aws --endpoint-url=http://localhost:4566 s3 ls

# Lista todos os objetos em um bucket S3
aws --endpoint-url=http://localhost:4566 s3 ls s3://meu-bucket --recursive

# Imprime o conteúdo de um arquivo específico no bucket S3
aws --endpoint-url=http://localhost:4566 s3 cp s3://meu-bucket/output.json -

# Parar o container 
docker stop localstack

# Reiniciar o container
docker start localstack


