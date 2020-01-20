<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Manutenção de Funcionarios</h1>
      </div>
      <div>
        <DataTable
          :value="funcionarios"
          :selection.sync="selectedFuncionarios1"
          :filters="filters"
          selectionMode="single"
          data-key="cpf.String"
          :paginator="true"
          :rows="10"
        >
          <!-- @row-select="onRowSelect" @row-unselect="onRowUnselect"> -->
          <template #header>
            <div style="text-align: right">
              <i class="pi pi-search" style="margin: 4px 4px 0px 0px;"></i>
              <InputText v-model="filters['global']" placeholder="Busca Global (F2)" size="50" />
            </div>
          </template>
          <Column field="foto" header="Foto">
            <template #body="slotProps"> 
              <img    
              :src="mostraFotoNoGrid(slotProps.data.funcid.String)"
              width="64px"
              /> 
            </template> 
          </Column>
          <Column field="funcnome.String" header="Nome Funcionário" filterMatchMode="startsWith"></Column>
          <Column field="cpf.String" mask="999.999.999-99" header="CPF" filterMatchMode="contains"></Column>
          <Column field="rg.String" header="RG" filterMatchMode="contains"></Column>
          <Column field="datanasc.String" header="Data de Nascimento" filterMatchMode="contains"></Column>
          <Column
            field="funcdatacontratacao.String"
            header="Data de Contratação"
            filterMatchMode="contains"
          ></Column>
          <Column
            field="funcdatadispensa.String"
            header="Data de Dispensa"
            filterMatchMode="contains"
          ></Column>
          <template #empty>Nenhum registro encontrado</template>
          <template #footer>
            <div style="text-align:center">
              <Button
                label="Cadastro"
                @click="chamacadastrodefuncionario"
                class="p-button-rounded"
              />
              <Button 
                label="Alterar" 
                @click="alteraDadosDeFuncionarios"
                class="p-button-rounded p-button-warning" 
              />
              
              <Button label="Demitir funcionário" class="p-button-rounded p-button-danger" />
            </div>
          </template>
        </DataTable>
      </div>
    </div>
  </div>
</template>

<script>
//import UsuariosService from "../../service/UsuariosService";
export default {
  data() {
    return {
      filters: {},
      fotoGrid: null,

      funcionarios: null,
      selectedFuncionarios1: null,
    };
  },
  funcionariosService: null,
  methods: {
    chamacadastrodefuncionario() {
      this.$router.push("/cadastrodefuncionario");
    },
    demitirfuncionario() {
      // this.$router.
      this.$router.push("/cadastrodefuncionario");
    },
    alteraDadosDeFuncionarios() {
      // this.$router.
      // this.$router.push("/cadastrodefuncionario");

      var idFuncionario = 3;

      localStorage.setItem('storeIdFuncionario', JSON.stringify(idFuncionario));

      this.$router.push("/cadastrodefuncionario");
    },
    mostraFotoNoGrid(idFuncionario){
      
      // const formData = new FormData();
      // formData.append("foto", this.selectedFile, this.selectedFile.name); Retirei o último parametro, se der pau, volto ele
      // console.log("Valor ID:", idFuncionario)
      // console.log("BaseUrl foto:", this.$httpBaseURL)
      

      return this.$httpBaseURL + '/rh/retornafotofuncionario/' + idFuncionario
      // formData.append("foto", idfuncionario);
      // this.$http
      //   .post("/rh/retornafotofuncionario", formData, { responseType: "blob" })
      //   .then(res => {
      //     let reader = new FileReader();
      //     reader.readAsDataURL(res.data);
      //     reader.onload = () => {
      //       this.fotoGrid = reader.result;
      //     };
      //   });
      //return this.fotoGrid;
    }
  },
  created() {
    // this.usuariosService = new UsuariosService();
  },
  mounted() {
    // this.usuariosService.getTodosUsuarios().then(data => (this.usuarios = data));

    //Limpa as Storages ao iniciar o formulário
    localStorage.removeItem('storeIdFuncionario');

    this.$http.get("/rh/listaTodosFuncionarios").then(res => {
      this.funcionarios = res.data.resposta;
      // console.log("Foto:" + this.funcionarios)
    });
    
  }
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>