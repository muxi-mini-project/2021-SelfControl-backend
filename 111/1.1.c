#include<stdio.h>
#include<math.h>

int haha(int num) {
	int i, half;
	if (num == 3){
		return 1;
    }else if (num%6!=1&&num%6!=5){
        return 0;
    }
	half=(int)sqrt(num);
	for(i=5;i<=half;i+=6){
		if(num%i==0||num%(i+2)==0){
			return 0;
        }
	}
	return 1;
}

void main(){
	int i,m,n,j,s[200000];
	scanf("%d",&m);
    s[0]=2;
    int k=1;
    for(i=3;i<1000000;i+=2){
        if (haha(i)){
            s[k]=i;
            k++;
        }
    }
    //for (i=0;i<1000;i++)printf("%d ",s[i]);
    for (i=1;i<100000;i++){
        n=m-s[i];
		for(j=i;s[j]<=n;j++){
            if(n==s[j]){
                printf("%d %d\n",s[i],n);
			    return;	
            }
        }
    }     
}