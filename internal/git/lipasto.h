#ifndef LIPASTO_H
#define LIPASTO_H

/*
 * Size of regular SHA-1 hash in git + 1
 * GIT_OID_SHA_HEXSIZE + 1
 */
#define OID_MAX_SIZE 41

struct bare_repository {
	char *path;
	char *name;
	char *owner;
	char *description;
	/*
	 * If hide is set to 1 it hides the repository from index. If ignore is
	 * set to 1 it hides the repository from index AND prevents it being
	 * accessed via direct path.
	 */
	int hide;
	int ignore;
};

#endif /* LIPASTO_H */
