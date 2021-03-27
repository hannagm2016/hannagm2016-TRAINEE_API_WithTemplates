<template>
 <div class="row py-lg-5">
                <div class="col-lg-6 col-md-8 mx-auto">
                    <form @submit="postData" method="post" >
                        <input type="text" name="title" placeholder="Title of your post" class="form-control"  v-model="posts.title"><br>
                        <textarea name="body" id="body"  placeholder="Body of your post" class="form-control"  v-model="posts.body"></textarea><br>
                        <button type="submit" class="btn btn-primary my-2">Save</button>
                    </form>
                </div>

            </div>

  </div>
</template>


<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)
export default {
    data (){
    return {
      posts:{
        id: null,
        userId: null,
        title: null,
        body:null
      }
    }
  },
  methods: {
     mess: function (message, event) {
     if (event) {
          event.preventDefault()
        }
          alert(message);
          window.location = '/';
    },
    postData(e){
           this.axios.post("http://localhost:8000/savePost", this.posts)
           .then((result)=>{
            console.warn(result)
           })
           e.preventDefault();

    }

  }
}
</script>


