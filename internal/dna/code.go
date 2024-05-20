package dna

var Dna = []string{"A", "C", "T", "G"}
var Rna = []string{"A", "C", "T", "U"}

var DnaComp = map[byte]byte{'A': 'T', 'C': 'G', 'G': 'C', 'T': 'A'}
var RnaComp = map[byte]byte{'A': 'U', 'C': 'G', 'G': 'C', 'U': 'A'}
