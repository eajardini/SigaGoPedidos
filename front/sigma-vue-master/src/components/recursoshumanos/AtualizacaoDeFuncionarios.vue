<template>
  <div class="p-grid p-fluid">
    <div class="p-col-12">
      <div class="card">
        <h1>Atualização dos Dados de Funcionários</h1>
      </div>
    </div>
    <TabView class="p-col-12 p-lg-12  ">
      <TabPanel header="Dados Pessoais" class="p-col-12 p-grid ">
        <div class="p-col-12 p-lg-9">
          <div class="card">
            <div class="p-grid">
              <div class="p-col-12 p-md-10">
                <p>Nome Funcionário</p>
                <InputText type="text" v-model="nomeFunc" />
              </div>
              <div class="p-col-12 p-md-3">
                <p>CPF</p>
                <!-- placeholder="999.999.999-99" -->
                <InputMask mask="999.999.999-99" v-model="CPFFunc"/>
              </div>
              <div class="p-col-12 p-md-3">
                <p>RG</p>
                <InputText type="text" v-model="RGFunc" />
              </div>
            </div>

            <div class="p-grid">
              <div class="p-col-12 p-md-2">
                <p>CEP</p>
                <InputMask mask="99999-999" v-model="CEPFunc" />
              </div>
              <div class="p-col-12 p-md-4">
                <p>Endereço</p>
                <InputText type="text" v-model="EnderFunc" />
              </div>
              <div class="p-col-12 p-md-3">
                <p>Cidade</p>
                <InputText type="text" v-model="CidadeFunc" />
              </div>
              <div class="p-col-12 p-md-1">
                <p>UF</p>                
                <InputMask mask="aa" v-model="UFFunc" />
              </div>
              <div class="p-col-12 p-md-3">
                <p>Estado</p>
                <InputText type="text" v-model="EstadoFunc" />
              </div>
              <div class="p-col-12 p-md-2">
                <p>Data Nascimento</p>
                <InputMask mask="99/99/9999"  v-model="DataNascFunc" />
              </div>
            </div>
          </div>
        </div>

        <div class="p-col-3 p-lg-3">
          <div class="card card-w-title">
            <h1>Foto do Funcionário</h1>
            <div class="p-grid">
              <div class="p-col-12 p-md-4">
                <img witdh="100px" height="100px" v-bind:src="fotoFuncionario" />
                <br />
                <input @change="EscolhiFoto" type="file" id="selectedFile" style="display: none;" />
                <input
                  type="button"
                  value="Escolha uma foto"
                  onclick="document.getElementById('selectedFile').click();"
                />
              </div>
            </div>
          </div>
        </div>
      </TabPanel>
      <TabPanel header="Dados Fucionais" >                         
          <div class="p-col-12 p-lg-12 p-grid p-fluid">            
          <div class="card">   
            <div class="p-grid">
              <div class="p-col-12 p-md-3">
                <p>Data Contratação</p>
                <InputMask mask="99/99/9999" v-model="DataContratacaoFunc" />
              </div>
              <div class="p-col-12 p-md-3">
                <p>Data Dispensa</p>
                <InputMask mask="99/99/9999" v-model="DataDispensaFunc" />
              </div>
              <div class="p-col-12 p-md-4">
                <p>Salário</p>
                 <money                        
                       class="form-input input-lg" 
                       style="text-align: right; height:31px; font-size: 16px" 
                       v-model="SalarioFunc" 
                       v-bind="money"                      
                      ></money>
              </div>
              
            </div>
          </div>
        </div>      
      </TabPanel>
    </TabView>

    <div class="p-col-12 p-lg-12">
      <div class="card">
        <div class="p-grid">
          <div class="p-col-12 p-md-2" style="text-align: center">
            <Button
              @click="salvarDados"
              label="Salvar os Dados Alterados"
              class="p-button-raised p-button-rounded"
            />
          </div>
         
          <div class="p-col-12 p-md-2" style="text-align: center">
            <Button
              @click="chamamanutencaodefuncionario"
              label="Fechar Formulário"
              class="p-button-rounded p-button-warning"
            />
          </div>
        </div>
      </div>
    </div>  
    <div class="p-col-12">
        
         
          <Dialog :header= "HeaderMensagemDoServidor" :visible.sync="displayMensagem" :style="{width: '50vw'}" :modal="true">
            <p>
              {{MensagemDoServidor}}
            </p>
            <template #footer>
              <Button label="Fechar Mensagem" icon="pi pi-check" @click="fechaMensagem" />              
            </template>
          </Dialog>         
        
    </div>
  </div>

</template>

<script>
//import UsuariosService from "../../service/UsuariosService";
export default {
  data() {
    return {
      filters: {},
      selectedFuncionarios1: null,
      funcid: null,
      nomeFunc: null,
      CPFFunc: null,
      RGFunc: null,
      CEPFunc: null,
      EnderFunc: null,
      CidadeFunc: null,
      UFFunc: null,
      EstadoFunc: null,
      fotoFuncionario: null,
      fotoFuncionarioParaAtualizar: null, //necessario pois a variav. fotoFuncionario contem a foto + as palavras image/jpg;base64,
      DataNascFunc:         null,
      DataContratacaoFunc:  null,
      DataDispensaFunc:     null,
      SalarioFunc: null,
      money: {
        decimal: ",",
        thousands: ".",
        prefix: "",
        suffix: "",
        precision: 2,
        masked: false
      },
      displayMensagem: false,
      MensagemDoServidor : "",
      HeaderMensagemDoServidor : "",
      DemitirFuncionario : false,
      idFuncionarioAlteracao : "",
      funcionarios: null,
    };
  },
  funcionariosService: null,
  methods: {
    alteraDados(){

    },
    chamamanutencaodefuncionario() {
      this.$router.push("/manutencaodefuncionario");
    },
    fechaMensagem(){
      this.displayMensagem = false;
      this.$router.push("/manutencaodefuncionario");

    },
    atribuiDadosAoFormulario(parFuncionario) {
      // console.log("[atribuiDadosAoFormulario] Valor do parFuncionário:" + parFuncionario);
        (this.nomeFunc = parFuncionario.nomeFunc.String),
        (this.CPFFunc = parFuncionario.CPFFunc.String),
        (this.RGFunc = parFuncionario.RGFunc.String),
        (this.CEPFunc = parFuncionario.CEPFunc.String),
        (this.EnderFunc = parFuncionario.EnderFunc.String),
        (this.CidadeFunc = parFuncionario.CidadeFunc.String),
        (this.UFFunc = parFuncionario.UFFunc.String),
        (this.EstadoFunc = parFuncionario.EstadoFunc.String),
        (this.DataNascFunc = parFuncionario.DataNascFunc.String),
        (this.DataContratacaoFunc = parFuncionario.DataContratacaoFunc.String ),
        (this.DataDispensaFunc = parFuncionario.DataDispensaFunc.String),
        (this.fotoFuncionario = this.mostraFotoFuncionarioNoForm(parFuncionario.fotoFuncionario)),  
        (this.fotoFuncionarioParaAtualizar = parFuncionario.fotoFuncionario),
        (this.SalarioFunc = parFuncionario.money.Float64);   
        console.log("parFuncionario.fotoFuncionario:" + parFuncionario.fotoFuncionario)    ;

    },
    salvarDados() {
      const formData = new FormData();
      formData.append("funcid", this.idFuncionarioAlteracao);
      formData.append("nomeFunc", this.nomeFunc);
      formData.append("CPFFunc", this.CPFFunc);
      formData.append("RGFunc", this.RGFunc);
      formData.append("CEPFunc", this.CEPFunc);
      formData.append("EnderFunc", this.EnderFunc);
      formData.append("CidadeFunc", this.CidadeFunc);
      formData.append("UFFunc", this.UFFunc);
      formData.append("EstadoFunc", this.EstadoFunc);
      formData.append("DataNascFunc", this.DataNascFunc);
      formData.append("DataContratacaoFunc", this.DataContratacaoFunc);
      formData.append("DataDispensaFunc", this.DataDispensaFunc);
      formData.append("SalarioFunc", this.SalarioFunc);          
      if (this.selectedFile !=null) {
        console.log("this.selectedFile:" + this.selectedFile);
         formData.append("foto", this.selectedFile); 
      } else {
        console.log("fotoFuncionarioParaAtualizar:" + this.fotoFuncionarioParaAtualizar);
         formData.append("foto", this.fotoFuncionarioParaAtualizar);
      }

      this.$http
        .post("/rh/atualizaFuncionarios", formData)
        .then(response => {
          this.MensagemDoServidor = response.data;
          this.HeaderMensagemDoServidor = "Sucesso";
          this.displayMensagem = true;
          // this.chamamanutencaodefuncionario();
        })
        .catch(error => {
          this.HeaderMensagemDoServidor = "Atenção!";
          this.MensagemDoServidor =
            "Problema ao atualizar o funcionário: " + error.response.data;
          this.displayMensagem = true;
        });
    },

    BuscaDadosParaAtualizacaoDeFuncionarios(idFuncionarioAlteracao){
      const formData = new FormData();

      
      formData.append("funcid", idFuncionarioAlteracao);      
      this.$http.post("/rh/buscaDadosFuncionariosParaAtualizar", formData)
      .then(res => {        
          this.funcionarios = res.data;
          this.atribuiDadosAoFormulario(this.funcionarios)
      }).catch(() => {
          this.HeaderMensagemDoServidor = "Atenção!"
          this.MensagemDoServidor = "Problema ao alterar o funcionário";
          this.displayMensagem = true;
        });
    },
    EscolhiFoto(event) {
      this.selectedFile = event.target.files[0];
      this.onUpload();
    },
    onUpload() {
      //https://www.youtube.com/watch?v=J84uiTfjbGA
      const formData = new FormData();
      // formData.append("foto", this.selectedFile, this.selectedFile.name); Retirei o último parametro, se der pau, volto ele
      formData.append("foto", this.selectedFile);
      this.$http
        .post("/rh/upLoadFotoFuncionario", formData, { responseType: "blob" })
        .then(res => {
          let reader = new FileReader();
          reader.readAsDataURL(res.data);
          reader.onload = () => {
            this.fotoFuncionario = reader.result;
          };
        });
    },

    mostraFotoFuncionarioNoForm(foto){       
      return "data:image/jpg;base64, " + foto;
      // return this.$httpBaseURL + '/rh/retornafotofuncionario/' + idFuncionario
    },

    previewFiles() {
      this.files = this.$refs.myFiles.files;
    }
  },
  created() {
    // this.usuariosService = new UsuariosService();
  },
  mounted() {
    // this.usuariosService.getTodosUsuarios().then(data => (this.usuarios = data));
    if (localStorage.getItem('storeIdFuncionario')) {
      try {
        this.idFuncionarioAlteracao = JSON.parse(localStorage.getItem('storeIdFuncionario'));
        console.log("[CafastgroDeFucionarios] Valor do ID Funcionário:" + this.idFuncionarioAlteracao)
        localStorage.removeItem('storeIdFuncionario');
        this.BuscaDadosParaAtualizacaoDeFuncionarios(this.idFuncionarioAlteracao);
      } catch(e) {
        localStorage.removeItem('storeIdFuncionario');
      }
    }
  }
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>