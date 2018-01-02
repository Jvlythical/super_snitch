# super_snitch

Super snitch is source code plagiarism detection service.

## Design

Super snitch is a super-powered plagiarism detection service.

A short-term goal of super snitch is to employ existing plagiarism detection
tools like MOSS and Deckard to help generate a super report of possible
plagiarism. A long-term goal is to find the best method of detecting plagiarism
given some context.

Users of plagiarism detection services typically upload a collection of files
and receive a report.  The report is essentially a sequence of triples where
each triple consist of a pair of files, file1 and file2, and a similarity score
between those files.  One obvious network optimization we add is the ability to
update an existing collection of files.

Super snitch will store the collection of files and call a plagiarism detection
engine.  Engines like MOSS generate their own reports, so Super Snitch will
simply return that.  Engines like Deckard do not appear to have web-based
reports so we have to support that.

## Dependencies

* go
* moss.pl 
  - The MOSS perl script.
  - This will be converted to a full Go library soon.

## Usage 


## API

