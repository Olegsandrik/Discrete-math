#include<iostream>
#include<vector>
#include<map>
#include<set>

using namespace std;

class autoMoore{
public:
    vector<vector<int>> input;
    map<pair<int, int>, int> num;
    set<pair<int, int>> P;
    vector<string> inputA;
    vector<string> outputA;
    void graph(){
        cout<<"digraph {\n    rankdir = LR\n";
        int i = 0;
        for(auto it : P){
            cout<<"    "<<i++<<" [label = \"("<<it.first<<","<<outputA[it.second]<<")\"]\n";
        }
        i = 0;
        for(auto it : P){
            for(int j = 0 ; j < input[num[it]].size() ; j++){
                cout<<"    "<<i<<" -> "<<input[num[it]][j]<<" [label = \""<<inputA[j]<<"\"]\n";
            }
            i++;
        }
        cout<<"}";
    }
};

class autoMealy{
public:
    int m, n, k;
    vector<vector<int>> input;
    vector<vector<int> > output;
    vector<string> inputA;
    vector<string> outputA;
    void newA(){
        cin>>m;
        inputA = vector<string>(m);
        for (int i = 0 ; i < m ; i++){
            cin>>inputA[i];
        }
        cin>>k;
        outputA = vector<string>(k);
        for(int i = 0 ; i < k ; i++){
            cin>>outputA[i];
        }
        cin>>n;
        input = vector<vector<int>>(n, vector<int>(m));
        output = vector<vector<int>>(n, vector<int>(m));
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
    autoMoore transformation(){
        autoMoore newA;
        set<pair<int, int>> NewP;
        for(int i = 0 ; i < n ; i++){
            for(int j = 0 ; j < m ; j++){
                NewP.insert(pair<int, int>(input[i][j], output[i][j]));
            }
        }
        map<pair<int, int>, int> num;
        int i = 0;
        for(auto it : NewP){
            num[it] = i;
            i++;
        }
        newA.num = num;
        newA.input = vector<vector<int>>(num.size(), vector<int>(m));
        for(auto it : NewP){
            for(i = 0 ; i < m ; i++){
                newA.input[num[it]][i] = num[pair<int, int>(input[it.first][i], output[it.first][i])];
            }
        }
        newA.inputA = inputA;
        newA.outputA = outputA;
        newA.P = NewP;
        return newA;
    }
};

int main(){
    autoMealy A;
    A.newA();
    autoMoore B = A.transformation();
    B.graph();
}
