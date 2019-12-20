<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Manutenção de Usuários</h1>
      </div>
      <div>
        <DataTable :value="cars" :selection.sync="selectedCar1" :filters="filters" selectionMode="single" dataKey="vin" 
            :paginator="true" :rows="10" @row-select="onRowSelect" @row-unselect="onRowUnselect">
          <template #header>
            <div style="text-align: right">
              <i class="pi pi-search" style="margin: 4px 4px 0px 0px;"></i>
              <InputText v-model="filters['global']" placeholder="Busca Global (F2)" size="50" />
            </div>
          </template>
          <Column field="vin" header="Login" filterMatchMode="startsWith"></Column>
          <Column field="year" header="Year" filterMatchMode="contains"></Column>
          <Column field="brand" header="Brand" filterMatchMode="equals"></Column>
          <Column field="color" header="Color" filterMatchMode="in"></Column>
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
import CarService from "../../service/CarService";
export default {
  data() {
    return {
      filters: {},
      brands: [
        { brand: "Audi", value: "Audi" },
        { brand: "BMW", value: "BMW" },
        { brand: "Fiat", value: "Fiat" },
        { brand: "Honda", value: "Honda" },
        { brand: "Jaguar", value: "Jaguar" },
        { brand: "Mercedes", value: "Mercedes" },
        { brand: "Renault", value: "Renault" },
        { brand: "Volkswagen", value: "Volkswagen" },
        { brand: "Volvo", value: "Volvo" }
      ],
      colors: [
        { name: "White", value: "White" },
        { name: "Green", value: "Green" },
        { name: "Silver", value: "Silver" },
        { name: "Black", value: "Black" },
        { name: "Red", value: "Red" },
        { name: "Maroon", value: "Maroon" },
        { name: "Brown", value: "Brown" },
        { name: "Orange", value: "Orange" },
        { name: "Blue", value: "Blue" }
      ],
      cars: null,
      selectedCar1: null,
    };
  },
  carService: null,
  created() {
    this.carService = new CarService();
  },
  mounted() {
    this.carService.getCarsLarge().then(data => (this.cars = data));
  }
};
</script>

<style scoped>
button {
  margin-right: 0.5em;
}
</style>