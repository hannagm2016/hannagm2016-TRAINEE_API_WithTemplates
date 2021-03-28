<template>
<div>

<main class="form-signin">
    <form @submit.prevent="registrateUser">
        <h1 class="h3 mb-3 fw-normal">Please Registrate</h1>
         <input v-model="User.Name" type="text"  id="inputName" class="form-control" placeholder="Name" required="" autofocus="">
        <input v-model="User.Email" type="email" id="inputEmail"  name="inputEmail" class="form-control" placeholder="Email address" required="" autofocus="">
         <input v-model="User.Password" type="password" id="inputPassword" name="inputPassword" class="form-control" placeholder="Password" required="">
        <input v-model="User.Password2" type="password" id="inputPassword2" name="inputPassword" class="form-control" placeholder="Confirm Password" required="">
        <button class="w-100 btn btn-lg btn-primary" type="submit">Registrate</button>
    </form>
    <div class="text-muted py-3">
        <a href="/#/authorization">Authorization</a>
    </div>
</main>
  </div>
</template>

<script>
import axios from 'axios'
export default {
   data: () => ({
      User: {
        Name: '',
        Email: '',
        Password: '123',
        Password2: '123'
      }
    }),
    methods: {
      registrateUser() {
      axios.post (`http://localhost:8000/registrationPost`, {
                     Name: this.User.Name,
                     UserName: this.User.Name,
                     Email: this.User.Email,
                     Password: this.User.Password,
                     //Password2: this.User.Password2,
                })
                 .then (response => {
                console.log (response)
               alert ("Registered: " + this.User.Name)
                 })
                  window.location = '/'
             }
}
}
</script>

<style>

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
  margin-top:100px;
 }
.form-signin .checkbox {
  font-weight: 400;
}
.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}
.form-signin .form-control:focus {
  z-index: 2;
}
.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

</style>
