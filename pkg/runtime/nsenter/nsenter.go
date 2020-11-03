package nsenter

/*
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *scr_pid;
	scr_pid = getenv("scr_pid");
	if (scr_pid) {
		//fprintf(stdout, "got scr_pid=%s\n", scr_pid);
	} else {
		//fprintf(stdout, "missing scr_pid env skip nsenter");
		return;
	}
	char *scr_cmd;
	scr_cmd = getenv("scr_cmd");
	if (scr_cmd) {
		//fprintf(stdout, "got scr_cmd=%s\n", scr_cmd);
	} else {
		//fprintf(stdout, "missing scr_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", scr_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(scr_cmd);
	exit(0);
	return;
}
*/
import "C"
