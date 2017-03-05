# Course 1 Finding Hidden Messages in DNA (Bioinformatics I)

## Week 1

### Where in the Genome does DNA replication begin?

You need to unwind DNA (two strands of DNA) and separate them.
You need to make sure that there are enough free-floating nucleotides so replication can complete because if you start replication and you have to stop in the middle because there is no energy left, you are dead.

Bacteria usually have circular chromosomes.
There must be some hidden messages in the genome that tell the cell "Start replication right here!".

A k-mer `Pattern` (a string of length `k`) is a most frequent word in text if no other k-mer is more frequent than `Pattern`.

Replication is performed by DNA polymerase, but to initiate replication, DNA polymerase needs a protein called DnaA.
DnaA binds to a short region (typically just a 9 nucleotide-long segment) within the replication origin that is known as a DnaA box.
A DnaA box is actually the hidden message that we are looking for, and the DnaA box actually tells the DnaA protein: bind right here!
And the DnaA protein wants to see multiple DnaA boxes to bind efficiently.

But if we don't know replication origin, what should we do?
Maybe we should look at all possible windows within a genome.
And if in a particular window, we find some frequent words appearing many times, it's a hint that maybe we found a replication origin.

A k-mer forms an (L, t)-clump inside genome if there is a short (length `L`) interval of genome in which it appears many (at least `t`) times.

#### Clump Finding Problem

Find patterns forming clumps in a string.
* Input - A string Genome and integers `k` (length of a pattern), `L` (window length), and `t` (number of patterns in a clump).
* Output - All k-mers forming (L, t)-clumps in Genome.

## Week 2

As strands unwind, they create two **replication forks**, which expand in both directions around the chromosome until the strands completely separate at the **replication terminus** (denoted *ter*).
The replication terminus is located roughly opposite to *ori* in the chromosome.

To start replication, a DNA polymerase needs a **primer**, a short complementary segment (shown in red above) that binds to the parent strand and jump starts the DNA polymerase.
After the strands start separating, each of the four DNA polymerases starts replication by adding nucleotides, beginning with the primer and proceeding around the chromosome from *ori* to *ter* in either the clockwise or counterclockwise direction.

However, nature has not yet equipped DNA polymerases with this ability, as they are **unidirectional**, meaning that they can only traverse a template strand of DNA in the 3' -> 5' direction, which is opposite from the 5’ -> 3’ direction of DNA.

DNA polymerase copies the template strand in the 3' -> 5' direction.
Note that it creates the daughter strand in the 5' -> 3' direction.

Skew(k): #G - #C for the first k nucleotides of Genome.
Skew diagram: Plot Skew(k) against k

