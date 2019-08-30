#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <fcnt1.h>
#include <dos.h>
#include <bios.h>
#include <io.h>
#include <graphics.h>
 
void initgraphs(int Vmode)
{
    union REGS inregs,outregs;
    inregs.h.ah=0;
    inregs.h.al=Vmode;
    int86(0x10,&inregs,&outregs);
}
 
void set_color_map(int i,int n,char *r,char *g,char *b)
{
    struct REGPACK regs; 
    char buff[256][3];
    int m;
    for(m=0;m<n;m++)
       {
          buff[m][0]=*r++>>2;
          buff[m][1]=*b++>>2;
          buff[m][2]=*g++>>2;
       }
    regs.r_ax=0x1012;
    regs.r_es=FP_SEG(buff);
    regs.r_dx=FP_OFF(buff);
    regs.r_bx=i;
    regs.r_cx=n;
    intr(0x10,®s);
}
 
int main()
{
    int i;
    FILE *fp;
    long vofs;
    int vseg;
    char far *ptr;
    char r[256],g[256],b[256],rev[256];
    int Vmode;
    char file_name[20];
 
    strcpy(file,"VGA256.bmp");
    Vmode=0x13;
    initgraphs(Vmode);
     
    if((fp=open(file_name,O_RDONLY))==1) exit();
    lseek(fp,54,0);
    for(i=0;i<256;i++)
       {
          read(fp,r+i,1);
          read(fp,b+i,1);
          read(fp,g+i,1);
          read(fp,rev+i,1);
       }
    set_color_map(0,256,g,r,b);
    vofs=0;
    lseek(fp,1078,0);
    ptr=(char far *)MK_FP(0xA000,vofs);
    read(fp,ptr,64000);
    close(fp);
    getch();
} 
