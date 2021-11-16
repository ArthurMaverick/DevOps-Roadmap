- pods é a menor unidade de informação que um sistema pode armazenar. 

- pods pode conter uma ou mais containers, que podem conter um ou mais volumes.

- os pods tem um endereço de rede, que é o endereço de rede que os containers podem usar para se comunicar com o sistema.
a.g: 

|  pod   | |container |
 10.0.0.1 :    8080

- 2 container ou mais nao podem ficar na mesma porta de um pod

- caso o container morra o pode morre junto e outro pod ira substituir o pod ira ser criado automaticamente e alocado no lugar do pod que morreu com um ip diferente

- pods podem compartilhar namespace de rede, IPC (Inter-Process Communication) e volumes

- pods do mesmo namespace podem se comunicar  entre si usando o mesmo namespace de rede e IPC