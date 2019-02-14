Massiminimo
===========

Descrizione
-----------

Scrivete un programma che, data una sequenza di al più 100 interi positivi
distinti nel flusso di ingresso, emetta nel flusso di uscita il massimo tra
essi, seguito dal minimo, seguito quindi dal massimo dei rimanenti e dal minimo
dei rimanenti e così via.

Ad esempio, se il flusso di ingresso contiene i numeri

    11 15 18
    16 12   10 19
    17
    20    14
    13

il programma emette

    10
    20
    11
    19
    12
    18
    13
    17
    14
    16
    15

nel flusso d'uscita; come si nota `10` è il più piccolo tra i valori in
ingresso, mentre `20` è il maggiore, quindi `11` e `19` sono rispettivamente il
più grande e più piccolo dei valori rimanenti, quindi lo sono `12` e `18` e così
via.


Vincoli
-------

I numeri nel flusso di ingresso sono separati da uno, o più, spazi bianchi e/o
segni di *a-capo*. Il flusso di ingresso potrebbe essere vuoto, ma non contiene
più di 100 valori e termina all'EOF. Il flusso di uscita deve contenere un
numero per linea.


Esempio
-------

Eseguendo `soluzione` e avendo `1 2 3` nel flusso di ingresso, il programma
emette

    1
    3
    2

nel flusso d'uscita.
