## fargate 
- fargate nao utiliza instancias ec2
- fargate cobra por tarefa e nao por instancias
- com fargate nao é necessario gerencia cluster

Oque é AWS Fargate no Amazon ECS?
  - O AWS Fargate o "gerenciado" da infraestrutura que roda containers . O fargate decide quais maquinas usar, tipo de server e como escalar.
  O AWS FArgatee permite executar containers sem a necessidade de gerenciar servidore e instancias do aws Ecs. Ou seja voce nao precisa se preocupar em escolher tipos de servidor , decidir quando o tamanho do cluster ou otimizar os recursos.


- como Fargate podemos especificar as quantidades de memori e cpu para os nosso conttainers
- toda a infraestrutura é provisionada e gerenciada automaticamnete pela aws, diferentemente do modelo ec2.