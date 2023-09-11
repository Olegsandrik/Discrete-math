#include<iostream>
#include<vector>

using namespace std;

class autoMealy{
public:
    int m, n, q0;
    vector<vector<int> > input;
    vector<vector<string> > output;
    void newA(){
        cin>>n>>m>>q0;
        input = vector<vector<int>>(n, vector<int>(m));
        output = vector<vector<string>>(n, vector<string>(m));
        for(int i = 0 ; i < n ; i++){
            for(int j = 0; j< m ; j++){
                cin>>input[i][j];
            }
        }
        for(int i = 0 ; i < n ; i++){
            for(int j = 0 ; j < m ; j++){
                cin>>output[i][j];
            }
        }
    }

    bool equal(autoMealy &O){
        if(O.m != m)
            return false;
        vector<vector<bool> > pos1(n, vector<bool>(O.n, false));
        return Equal(O, pos1, q0, O.q0);
    }

    bool Equal(autoMealy &O, vector<vector<bool> > &pos1, int q, int Oq){
        pos1[q][Oq] = true;
        for(int i = 0 ; i < m ; i++){
            if(output[q][i] != O.output[Oq][i]){
                return false;
            }
            if(!pos1[input[q][i]][O.input[Oq][i]]){
                if(!Equal(O, pos1, input[q][i], O.input[Oq][i])){
                    return false;
                }
            }
        }
        return true;
    }
};

int main(){
    autoMealy A1, A2;
    A1.newA();
    A2.newA();
    if(A1.equal(A2)){
        cout<<"EQUAL";
    }else{
        cout<<"NOT EQUAL";
    }
}
