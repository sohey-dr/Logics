use strict;
use warnings;

fizzbuzz();

sub fizzbuzz {
  my $i = 1;
  while ($i <= 100) {
    if ($i % 15 == 0) {
      print "FizzBuzz\n";
    } elsif ($i % 3 == 0) {
      print "Fizz\n";
    } elsif ($i % 5 == 0) {
      print "Buzz\n";
    } else {
      print "$i\n";
    }
    $i++;
  }
}