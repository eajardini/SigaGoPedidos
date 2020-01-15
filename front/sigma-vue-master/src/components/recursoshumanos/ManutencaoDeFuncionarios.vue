<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Manutenção de Funcionarios</h1>
      </div>
      <div>
        <DataTable :value="funcionarios" :selection.sync="selectedFuncionarios1" :filters="filters" selectionMode="single" dataKey="cpf.String" 
            :paginator="true" :rows="10"> <!-- @row-select="onRowSelect" @row-unselect="onRowUnselect"> -->
          <template #header>
            <div style="text-align: right">
              <i class="pi pi-search" style="margin: 4px 4px 0px 0px;"></i>
              <InputText v-model="filters['global']" placeholder="Busca Global (F2)" size="50" />
            </div>
          </template>
          <Column field="funcnome.String" header="Nome Funcionário" filterMatchMode="startsWith"></Column>
          <Column field="cpf.String" mask="999.999.999-99" header="CPF" filterMatchMode="contains"></Column>              
          <Column field="rg.String" header="RG" filterMatchMode="contains"></Column>                             
          <Column field="datanasc.String" header="Data de Nascimento" filterMatchMode="contains"></Column>                   
          <Column field="funcdatacontratacao.String" header="Data de Contratação" filterMatchMode="contains"></Column>                   
          <Column field="funcdatadispensa.String" header="Data de Dispensa" filterMatchMode="contains"></Column>                         
          <template #empty>Nenhum registro encontrado</template>
          <template #footer>
            <div style="text-align:center">
              <Button label="Cadastro (F3)" @click="chamacadastrodefuncionario" class="p-button-rounded" />
              <Button label="Alterar(F4)" class="p-button-rounded p-button-warning" />
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
     
      funcionarios: null,
      selectedFuncionarios1: null,
    };
  },
  funcionariosService: null,
  methods: {
    chamacadastrodefuncionario () {
      this.$router.push('/cadastrodefuncionario')
    },

  },
  created() {
   // this.usuariosService = new UsuariosService();
  },
  mounted() {
    // this.usuariosService.getTodosUsuarios().then(data => (this.usuarios = data));
    this.$http.get("/rh/listaTodosFuncionarios").then(res => {
                    this.funcionarios = res.data.resposta});            
  }
  
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>