/*
Package twrap provides supporting functions for printing wrapped text. It
allows you to print indented text that will fit within a supplied maximum
line length. It is also possible to specify a minimum amount of text to print
on each line in case the indent is too great.

You should first of all create a TWConf (use the NewTWConf function - it will
report invalid parameters). Then you can call the various Wrap...(...)
methods on it to print the text.
*/
package twrap
