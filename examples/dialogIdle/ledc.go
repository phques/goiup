package main

/*
#include <stdlib.h>
#include <stdarg.h>
#include <iup/iup.h>
*/
import "C"

/*

int ledLoad(void)
{
static Ihandle* named[      6 ];

  named[0] = IupSetAtt( "mainDialog", IupDialog(
    IupSetAtt( NULL, IupVbox(
      IupSetAtt( NULL, IupFrame(
        IupSetAtt( NULL, IupVbox(
          named[1] = IupSetAtt( "localRoot", IupText( "NULL" ),
            "EXPAND", "HORIZONTAL",
            "MARGIN", "5", NULL ),
        NULL),
          "NMARGIN", "5x5",
          "NGAP", "5x5",
          "EXPAND", "YES", NULL )
      ),
        "TITLE", "Local Root", NULL ),
      IupSetAtt( NULL, IupFrame(
        IupSetAtt( NULL, IupVbox(
          IupVbox(
            IupLabel( "Root" ),
            named[2] = IupSetAtt( "destinationRoot", IupList( "NULL" ),
              "1", "aa",
              "2", "bb",
              "3", "cc",
              "DROPDOWN", "YES",
              "EXPAND", "HORIZONTAL", NULL ),
          NULL),
          IupVbox(
            IupLabel( "Files" ),
            named[3] = IupSetAtt( "Files", IupList( "null" ),
              "EXPAND", "YES", NULL ),
          NULL),
        NULL),
          "NMARGIN", "5x5",
          "NGAP", "5x5",
          "EXPAND", "YES", NULL )
      ),
        "TITLE", "Destination", NULL ),
      IupSetAtt( NULL, IupHbox(
        IupFill(),
        named[4] = IupSetAtt( "pushButton", IupButton( "Push", "onPushButton" ), NULL ),
      NULL),
        "EXPAND", "HORIZONTAL",
        "SIZE", "150", NULL ),
    NULL),
      "NMARGIN", "5x5",
      "NGAP", "5x5",
      "EXPAND", "YES", NULL )
  ),
    "TITLE", "Android Push",
    "MARGINS", "5x5", NULL );
}


*/
import "C"
