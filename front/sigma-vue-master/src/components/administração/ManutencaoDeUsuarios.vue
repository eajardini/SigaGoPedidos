<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Manutenção de Usuários</h1>
      </div>
      <div >
        <DataTable :value="usuarios" :selection.sync="selectedUsuarios1" :filters="filters" selectionMode="single" dataKey="vin" 
            :paginator="true" :rows="10" > <!--@row-select="onRowSelect" @row-unselect="onRowUnselect">  -->
          <template #header>
            <div style="text-align: right">
              <i class="pi pi-search" style="margin: 4px 4px 0px 0px;"></i>
              <InputText ref="busca" v-model="filters['global']" placeholder="Busca Global (F2)" size="50" />
            </div>
          </template>
          <Column field="funcnome" header="Nome Funcionário" filterMatchMode="startsWith"></Column>
          <Column field="login" header="Login" filterMatchMode="contains"></Column>          
          <Column field="datacriacao" header="Data da Criação" filterMatchMode="contains"></Column> 
          <Column field="datavalidade" header="Data da Validade" filterMatchMode="contains"></Column>                   
          <Column field="userbloqueado" header="Usuário Bloqueado ?" filterMatchMode="contains"></Column>                   
          <template #empty>Nenhum registro encontrado</template>
          <template #footer>
            <div style="text-align:center">
              <Button label="Cadastro (F3)" class="p-button-rounded" />
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
     
      usuarios: null,
      selectedUsuarios1: null,
    };
  },
  usuariosService: null,
  created() {
   // this.usuariosService = new UsuariosService();
  },
  mounted() {
    // this.usuariosService.getTodosUsuarios().then(data => (this.usuarios = data));
    this.$http.get("/listaTodosUsuarios").then(res => {
                    this.usuarios = res.data.resposta});            
  },
 methods: {
    pegaFocus: function() {
      // Note, you need to add a ref="search" attribute to your input.
      alert("Borodin");
      this.$refs.busca.focus();
    },
 },
  
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>