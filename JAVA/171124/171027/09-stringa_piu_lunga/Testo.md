Stringa più lunga
=================

Descrizione
-----------

Scrivete un programma che, data una sequenza di stringhe nel flusso di ingresso,
emetta nel flusso d'uscita la stringa di lunghezza maggiore tra esse, emettendo
la prima che incontra tra esse qualora ci fossero più stringhe di lunghezza
massima.

Ad esempio, su input `la mia mamma che pasta buona mi fa`, il programma emette
`mamma` dato che è la prima stringa lunga 5 nel flusso e tutte le stringhe del
flusso hanno lunghezza minore o uguale a 5.


Vincoli
-------

Il flusso d'ingresso contiene una o più stringhe, separate da un qualunque
carattere *white-space*, non necessariamente spazio o a-capo.


Esempio
-------

Eseguendo

    soluzione

e avendo nel flusso di ingresso

    che
    essere bello

il programma emette nel flusso di uscita

    essere

che infatti è la parola più lunga.


Suggerimenti
------------

Lo `Scanner` è in grado di segnalare che non ci sono più stringhe nel flusso di
ingresso restituendo il valore `false` alla chiamata di `hasNext()`; un ciclo
che analizzi in sequenza tutte le stringhe del flusso di ingresso ha pertanto la
forma seguente

    while (sc.hasNext()) {
        String str = sc.next();
        ...
    }

dove `sc` è un oggetto di tipo `Scanner` opportunamente istanziato.
