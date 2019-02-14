# Lezioni del laboratorio di "Programmazione" del corso di "Informatica Musicale"

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)
[![License: CC BY-SA 4.0](https://i.creativecommons.org/l/by-sa/4.0/80x15.png)](http://creativecommons.org/licenses/by-sa/4.0/)

Questo repository contiene i testi degli esercizi del [laboratorio](https://santini.di.unimi.it/d/labprog/) dell'insegnamento di [Programmazione](http://www.ccdinfmi.unimi.it/it/corsiDiStudio/2017/F3Xof2/default/F3X-36/index.html) del corso di laurea in [Informatica Musicale](http://www.ccdinfmi.unimi.it/it/corsiDiStudio/2017/F3Xof2/), suddivisi in una directory per ciascuna lezione.

Il contenuto relativo ad una lezione coincide sostanzialmente con quanto descritto nella guida "[Istruzioni
per lo svolgimento degli esercizi del laboratorio di
programmazione](http://reaper.srv.di.unimi.it/manuale.html)" a cui si rimanda per maggiori informazioni.

## In laboratorio

Durante lo svolgimento delle lezioni, il materiale contenuto in questo
repository sarà reso disponibile agli studenti tramite la [Scythe Suite](https://github.com/scythe-suite),
una collezione di strumenti che consentono la somministrazione, raccolta e
valutazione di esercizi di programmazione.

Per configurare il proprio account, al primo utilizzo, ogni studente deve
impartire il seguente comando alla shell

    curl -sL https://git.io/labprog-tools | bash
    source $HOME/.bashrc

o alternativamente (qualora il comando `curl` non fosse disponibile), il comando

    wget -qO- https://git.io/labprog-tools  | bash
    source $HOME/.bashrc

Il docente provvederà quindi a dare le ulteriori necessarie istruzioni durante
lo svolgimento delle lezioni.

## A casa

Chi intendesse utilizzare questo materiale al di fuori del laboratorio (ad
esempio per esercitarsi a casa, o per prepararsi all'esame) può scaricare un
[archivio compresso](https://github.com/mapio/labprog-infomus/archive/master.zip)
di questo repository, oppure ottenere un *clone* locale con il comando `git clone
git@github.com:mapio/labprog-infomus.git`; in questo caso si osserva che:

- i comandi `firma`, `aggiorna` e `consegna` non sono definiti, non è infatti prevista
  una modalità di autenticazione e scambio di file nel caso di utilizzo al di fuori del
  laboratorio;

- non è presente il file `LEGGIMI.md`, sostituito da questo file;

- la directory `bin/`, coi comandi `soluzione` e `verifica` che possono essere
  utilizzati anche a casa, è presente in una unica copia, al primo livello di
  questo repository; per poter usare tali comandi è necessario configurare il
  *path* del proprio interprete di comandi.

Se si usa l'interprete di comandi `bash`, per poter eseguire i comandi senza
indicarne il percorso completo, è sufficiente aggiornare la variabile d'ambiente
`PATH` semplicemente portandosi nella directory in cui è contenuto questo file
`README.md` ed impartire il comando

	export PATH="$(pwd)/bin:$PATH"

## Nota di copyright

I testi ed i sorgenti contenuti in questo repository sono protetti dalle leggi
sul copyright e dalle disposizioni dei trattati internazionali; in particolare
ai testi si applica la [Creative Commons Attribution-ShareAlike 4.0
International License](http://creativecommons.org/licenses/by-sa/4.0/) ed ai
sorgenti la [GNU General Public License
3](https://www.gnu.org/licenses/gpl-3.0.html).

L'informazione qui contenuta è ritenuta essere accurata alla data della
pubblicazione. Essa è fornita per scopi esclusivamente didattici e non per
essere utilizzata in progetti di impianti, prodotti, ecc. Tale informazione è
soggetta a cambiamenti senza preavviso. L'autore non si assume alcuna
responsabilità per il contenuto di queste pagine (ivi incluse, ma non
limitatamente a, la correttezza, completezza, applicabilità ed aggiornamento
dell'informazione). In ogni caso non può essere dichiarata conformità
all'informazione in oggetto.

In ogni caso questa nota di copyright non deve mai essere rimossa e deve essere
riportata anche in utilizzi parziali.

© Massimo Santini 2017
