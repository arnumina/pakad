/*
#######
##                    __           __
##         ___  ___ _/ /_____ ____/ /
##        / _ \/ _ `/  '_/ _ `/ _  /
##       / .__/\_,_/_/\_\\_,_/\_,_/
##      /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package component

// Logger AFAIRE
type Logger interface {
	Clone(rnr string) Logger
	Trace(msg string, ctx ...interface{})
	Debug(msg string, ctx ...interface{})
	Info(msg string, ctx ...interface{})
	Notice(msg string, ctx ...interface{})
	Warning(msg string, ctx ...interface{})
	Error(msg string, ctx ...interface{})
	Critical(msg string, ctx ...interface{})
	Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
