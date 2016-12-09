#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../vendor/zopflipng_lib.h"

void setfilters(CZopfliPNGOptions* png_options, int size, int* filters) {
    int* ptr = malloc(sizeof(int*)*size);
    if(ptr == NULL) {
    	return;
    }
    for (int i = 0; i < size; ++i) {
        ptr[i] = filters[i];
    }
    png_options->filter_strategies = (enum ZopfliPNGFilterStrategy*)ptr;
    png_options->num_filter_strategies = size;
    png_options->auto_filter_strategy = 1;
}

void setchunks(CZopfliPNGOptions* png_options, int size, char** chunks) {
    char** ptr = malloc(sizeof(char*)*size);
    if(ptr == NULL) {
    	return;
    }
    for (int i = 0; i < size; ++i) {
        ptr[i] = malloc(sizeof(char*)*10);
        strcpy(ptr[i], chunks[i]);
    }
    png_options->keepchunks = ptr;
    png_options->num_keepchunks = size;
}
void freeopts(CZopfliPNGOptions* png_options) {
    if (png_options->num_filter_strategies > 0 || png_options->filter_strategies != NULL) {
        free(png_options->filter_strategies);
        png_options->num_filter_strategies = 0;
        png_options->auto_filter_strategy = 0;
    }
    if (png_options->num_keepchunks > 0 || png_options->keepchunks != NULL) {
        for (int i = 0; i < png_options->num_keepchunks; ++i) {
            free(png_options->keepchunks[i]);
        }
        free(png_options->keepchunks);
        png_options->num_keepchunks = 0;
    }
}