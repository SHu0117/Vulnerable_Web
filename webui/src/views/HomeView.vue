<script>
import SuccessMsg from '../components/SuccessMsg.vue';
export default {
	components:{SuccessMsg},
	data: function() {
		return {
			errormsg: null,
			username: localStorage.getItem('username'),
			requesterID: localStorage.getItem('requesterID'),	
			succmsg: localStorage.getItem('succMsg'),		
			clear: null,		
		}
	},
	methods: {
		async refresh() {
		},
		async doLogout() {
			localStorage.removeItem("requesterID")
			localStorage.removeItem("username")
			localStorage.removeItem("succMsg")
			this.$router.push({ path: '/' })
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div>
			<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2 my-auto"><strong>Home page - Welcome back to WASA-PHOTO</strong></h1>
				<div class="d-flex align-items-center">
					<!-- Grouped Buttons and Input for Alignment -->					
					<i class="fas fa-sign-out-alt icon-spacing" @click="doLogout()" aria-label="Logout"></i>					
				</div>
			</div>
		</div>
	</div>

    <SuccessMsg v-if="this.succmsg" :msg="this.succmsg"></SuccessMsg>
</template>




<style>

.btn {
    height: 40px;
    padding: 8px 12px;
    font-size: 1rem; 
    border-radius: 4px;
}

.btn-custom {
    width: 100px;  
}

.custom-margin {
	margin-bottom: 30px; 
}

</style>

