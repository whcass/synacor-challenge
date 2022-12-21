use strict;
use warnings;

use List::Permutor;

my $permutor = List::Permutor->new( 2,7,3,9,5);
while ( my @permutation = $permutor->next() ) {
    my $answer = $permutation[0] + $permutation[1] * $permutation[2]**2 + $permutation[3]**3 - $permutation[4];
    if($answer == 399){
        print "FOUND IT! - @permutation\n";
    }
    # print "$answer\n";
}