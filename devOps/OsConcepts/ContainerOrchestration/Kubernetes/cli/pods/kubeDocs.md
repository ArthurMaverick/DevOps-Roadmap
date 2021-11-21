# Criando Objetos

kubectl apply -f ./my-manyfest.yml            '# criar recurso(s)
kubectl apply -f ./my1.yml -f ./my2.ymal      '# criar recursos a partir de varios arquivos
kubectl apply -f ./dir                        '# criar recursos e todos arquivos do diretorio
kubectl apply -f https://git.io/vPieo         '# criar recursos a partir de URL
kubectl create deployment nginx --image-nginx '# iniciar uma unica instancia do nginx
kubectl explay pods, svc                      '# Obtenha a documentação de manifesto do pod


<!-- create pods -->
kubectl **run** nginx-prod --image=nginx:latest
kubectl **run** nginx-prod --image=nginx:latest **--watch**  <!-- assite mudanças-->
 
<!--  describe pods -->
kubectl **describe** pod nginx-prod 


<!-- edit pods -->
kubectl **edit** pod nginx-prod

<!-- delete pods -->
kubectl **delete** -f ./manifest.yml
kubectl **delete** pods --all
kubectl **delete** svc --all

<!-- get -->
kubectl **get** pods
kubectl **get** pods -o wide  <!-- infos limpas do pod-->

<!-- exec -->
kubectl **exec** -it nginx-prod -- bash <!-- atach terminal da maquina ao terminal do container -->