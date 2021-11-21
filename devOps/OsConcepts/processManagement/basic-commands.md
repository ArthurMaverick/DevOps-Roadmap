# fg
  
  - Você pode usar o comando “fg” para continuar um programa que foi interrompido e trazê-lo para o primeiro plano.

  ```bash
  $ fg <PID>
  ```

  # top

  - Você pode usar o comando “top” para ver o estado atual de todos os processos.

  # ps ux

  - Você pode usar o comando “ps ux” para ver de todos os processos estaticamente. as informações exibidas são diferentes do top command 

  # ps PID

  - Você pode usar o comando “ps PID” para ver o estado atual de um processo específico.

  # kill PID

  - Você pode usar o comando “kill PID” para matar um processo específico.

  # killall

  - Você pode usar o comando “killall” para matar todos os processos.

  # killall -9

  - Você pode usar o comando “killall -9” para matar todos os processos, exceto os que estão em background.

  # nice -n nome fo processo 'Nice value'

    - O Linux pode executar muitos processos ao mesmo tempo, o que pode diminuir a velocidade de alguns processos de alta prioridade e resultar em baixo desempenho.

    Para evitar isso, você pode dizer à sua máquina para priorizar os processos de acordo com seus requisitos.

    Essa prioridade é chamada de Niceness no Linux e tem um valor entre -20 a 19. Quanto menor o índice de Niceness, maior será a prioridade dada a essa tarefa.

    O valor padrão de todos os processos é 0.

      Para iniciar um processo com um valor niceness diferente do valor padrão, use a seguinte sintaxe

  # renice 'valor' -p 'PID'

    - Para alterar o Niceness, você pode usar o comando 'top' para determinar o PID (id do processo) e seu valor Nice. Posteriormente, use o comando renice para alterar o valor.

    Vamos entender isso por um exemplo.


  # df 

  - Você pode usar o comando “df” para ver o espaço de disco disponível em sua máquina.

  # df -h

  - Você pode usar o comando “df -h” para ver o espaço de disco disponível em sua máquina, em MB.

  # free -h

  - Este comando mostra a memória livre e usada (RAM) no sistema Linux, em MB.