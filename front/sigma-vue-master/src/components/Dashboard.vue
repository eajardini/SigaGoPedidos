 <template>
<div class="p-grid p-fluid dashboard">
	<div class="p-col-12 p-lg-4">
		<div class="card summary">
			<span class="title">Usuários</span>
			<span class="detail">Número de visitas</span>
			<span class="count visitors">50</span>
		</div>
	</div>


</div>
	
	
	

</template>

<script>

import EventService from '../service/EventService';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';

export default {
	data() {
		return {
			tasksCheckbox: [],
			dropdownCities: [
				{name: 'New York', code: 'NY'},
				{name: 'Rome', code: 'RM'},
				{name: 'London', code: 'LDN'},
				{name: 'Istanbul', code: 'IST'},
				{name: 'Paris', code: 'PRS'}
			],
			dropdownCity: null,
			options: {
				plugins:[dayGridPlugin, timeGridPlugin, interactionPlugin],
				defaultDate: '2019-01-01',
				header: {
					left: 'prev,next',
					center: 'title',
					right: 'dayGridMonth,timeGridWeek,timeGridDay'
				},
				editable: true
			},
			events: null,
			dataTableCars: null,
			dataTableSelectedCar: null,
			lineData: {
				labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
				datasets: [
					{
						"label": 'First Dataset',
						data: [65, 59, 80, 81, 56, 55, 40],
						fill: false,
						backgroundColor: '#2f4860',
						borderColor: '#2f4860'
					},
					{
						"label": 'Second Dataset',
						data: [28, 48, 40, 19, 86, 27, 90],
						fill: false,
						backgroundColor: '#00bb7e',
						borderColor: '#00bb7e'
					}
				]
			},
		}
	},
	
	eventService: null,
	created() {
		
		this.eventService = new EventService();
	},
	mounted() {
		// this.carService.getCarsSmall().then(data => this.dataTableCars = data);
		// this.eventService.getEvents().then(data => this.events = data);
	},
}
</script>

<style scoped>

</style>