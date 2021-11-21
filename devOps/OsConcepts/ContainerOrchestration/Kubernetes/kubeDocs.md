# Criando Objetos

kubectl apply -f ./my-manyfest.yml            '# criar recurso(s)
kubectl apply -f ./my1.yml -f ./my2.ymal      '# criar recursos a partir de varios arquivos
kubectl apply -f ./dir                        '# criar recursos e todos arquivos do diretorio
kubectl apply -f https://git.io/vPieo         '# criar recursos a partir de URL
kubectl create deployment nginx --image-nginx '# iniciar uma unica instancia do nginx
kubectl explay pods, svc                      '# Obtenha a documentação de manifesto do pod
<<<<<<< HEAD

<!-- create pods -->
kubectl **run** nginx-prod --image=nginx:latest
                  <!-- --watch  asiste mudanças-->
 
<!-- describe pods -->
kubectl **describe** pod nginx-prod


<!-- edit pods -->
kubectl **edit** pod nginx-prod

