// Code generated by rgnonomic/rgo; DO NOT EDIT.

#include "_cgo_export.h"

void R_warning(char* s) {
	warning(s);
}

void R_error(char* s) {
	error(s);
}

// TODO(kortschak): Only emit these when needed:
// Needed for unpacking SEXP character.
GoString R_gostring(SEXP x, R_xlen_t i) {
	SEXP _s = STRING_ELT(x, i);
	GoString s = {(char*)CHAR(_s), STDVEC_LENGTH(_s)};
	return s;
}

// Needed for getting list elements by name.
int getListElementIndex(SEXP list, const char *str) {
	int index = -1;
	SEXP names = getAttrib(list, R_NamesSymbol);
	if (!isString(names)) {
		return index;
	}
	for (int i = 0; i < length(list); i++) {
		if (strcmp(CHAR(STRING_ELT(names, i)), str) == 0) {
			index = i;
			break;
		}
	}
	return index;
}

SEXP print_sexp(SEXP p) {
	return Wrapped_PrintSEXP(p);
}
