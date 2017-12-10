/*
** client.c -- a stream socket client demo
Based on Beej's public domain example at http://beej.us/guide/bgnet/output/html/multipage/clientserver.html#simpleclient

Compile with gcc monitor.c -Os -flto -lWs2_32
*/


#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>



#ifdef _WIN32
  #ifndef _WIN32_WINNT
    #define _WIN32_WINNT 0x0501  /* Windows XP. */
  #endif
#else
  /* Assume that any non-Windows platform uses POSIX-style sockets instead. */
  #include <sys/socket.h>
  #include <unistd.h> /* Needed for close() */

#endif

#include "../support/udp/json-parser/json.h"
#include "../support/udp/json-parser/json.c"


typedef struct {
  char *Selector;
  char *Arg;
} Message;

Message mess;
char * str;

#define MAXDATASIZE 10000 // max number of bytes we can get at once


static void print_depth_shift(int depth)
{
        int j;
        for (j=0; j < depth; j++) {
                printf(" ");
        }
}

static void process_value(json_value* value, int depth);

static void process_object(json_value* value, int depth)
{
        int length, x;
        if (value == NULL) {
                return;
        }
        length = value->u.object.length;
        for (x = 0; x < length; x++) {
                print_depth_shift(depth);
                //printf("object[%d].name = %s\n", x, value->u.object.values[x].name);
                process_value(value->u.object.values[x].value, depth+1);
                if(strcmp(value->u.object.values[x].name, "Selector")==0) {
                  mess.Selector = str;
                }
                if(strcmp(value->u.object.values[x].name, "Arg")==0) {
                  mess.Arg = str;
                }
        }
}

static void process_array(json_value* value, int depth)
{
        int length, x;
        if (value == NULL) {
                return;
        }
        length = value->u.array.length;
        printf("array\n");
        for (x = 0; x < length; x++) {
                process_value(value->u.array.values[x], depth);
        }
}

static void process_value(json_value* value, int depth)
{
        int j;
        if (value == NULL) {
                return;
        }
        if (value->type != json_object) {
                print_depth_shift(depth);
        }
        switch (value->type) {
                case json_none:
                        printf("none\n");
                        break;
                case json_object:
                        process_object(value, depth+1);
                        break;
                case json_array:
                        process_array(value, depth+1);
                        break;
                case json_integer:
                        printf("int: %10" PRId64 "\n", value->u.integer);
                        break;
                case json_double:
                        printf("double: %f\n", value->u.dbl);
                        break;
                case json_string:
                        //printf("string: %s\n", value->u.string.ptr);
                        str = value->u.string.ptr;
                        break;
                case json_boolean:
                        printf("bool: %d\n", value->u.boolean);
                        break;
                case json_null:
                        printf("null: \n");
                        break;
        }
}

int main(int argc, char *argv[]) {
    int numbytes;
    char buf[MAXDATASIZE];
    if (argc != 2) {
        fprintf(stderr,"use: monitor hostname:port\n");
		fprintf(stderr,"or\n");
		fprintf(stderr,"use: monitor pipes\n");
        exit(1);
    }

    for(;;) {
		gets(buf);

		buf[numbytes] = '\0';

		//printf("client: received '%s'\n",buf);
		json_char* json = (json_char*)buf;

        json_value* value = json_parse(json,numbytes);
        if (value == NULL) {
			fprintf(stderr, "Invalid message\n");
        } else {
        process_value(value, 0);
        fprintf(stderr, "%s:%s\n", mess.Selector, mess.Arg);
      }
	}
    return 0;
}
