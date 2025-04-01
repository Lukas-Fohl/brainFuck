#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <limits.h>

#define FILE_SIZE_MAX 1024
#define CELL_MAX 1024
#define CALL_BACK_MAX 32

int main(int argn, char * argv[]){
    FILE * f;
    f = fopen(argv[1],"r");
    char * out = malloc(FILE_SIZE_MAX*sizeof(char));
    fscanf(f,"%s",out);

    short cells[CELL_MAX] = {0};
    short* cellPtr = cells;

    short lastOpen[CALL_BACK_MAX]; 
    for(int i = 0; i < CALL_BACK_MAX; i++){
        lastOpen[i] = -1;
    }

    for(int i = 0; i < FILE_SIZE_MAX; i++){
        if(out[i] == '\0' || out[i] == EOF){
            break;
        }
        //printf("%c",out[i]);
        switch(out[i]){
            case '>':
                if(cellPtr >= cells+CELL_MAX){
                    printf("ERROR: out of bound");
                }
                cellPtr++;
                break;
            case '<':
                if(cellPtr <= cells){
                    printf("ERROR: out of bound");
                }
                cellPtr--;
                break;
            case '+':
                if((*cellPtr) == SHRT_MAX){
                    printf("ERROR: value overflow");
                }
                (*cellPtr)++;
                break;
            case '-':
                if((*cellPtr) == 0){
                    printf("ERROR: value underflow");
                }
                (*cellPtr)--;
                break;
            case '.':
                printf("%c",(*cellPtr));
                break;
            case '[':
                for(int j = 0; j < CALL_BACK_MAX; j++){
                    if(lastOpen[j] == -1){
                        lastOpen[j] = i;
                        break;
                    }
                }
                break;
            case ']':
                if((*cellPtr) <= 0){
                    (*cellPtr) = 0;
                    for(int j = 0; j < CALL_BACK_MAX; j++){
                        if(lastOpen[j] != -1){
                            lastOpen[j] = -1;
                            break;
                        }
                    }
                    break;
                }
                for(int j = 0; j < CALL_BACK_MAX; j++){
                    if(lastOpen[j] != -1){
                        i = lastOpen[j];
                        break;
                    }else if(j == CALL_BACK_MAX-1){
                        printf("ERROR: loop depth overflow (>%d)",CALL_BACK_MAX);
                    }
                }
                break;
            case ',':
                scanf("%c",cellPtr);
                break;
            default:
                break;
        }
    }
    return 0;
}
