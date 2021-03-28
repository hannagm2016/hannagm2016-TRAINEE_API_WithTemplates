<template>
  <div class="Index">
        <h5  v-if="posts.IsAuthorized" class="text-muted">Hello {{posts.Cust.UserName}}</h5>

          <section class="py-5 text-center container">
              <div class="row py-lg-5">
                  <div class="col-lg-6 col-md-8 mx-auto">
                      <h1 class="fw-light">All posts</h1>
                      <p class="lead text-muted">Only authorized users can leave coments and add or change posts</p>
                      <p>
                          <span v-if="posts.IsAuthorized">
                              <a href="/#/create" class="btn btn-secondary my-2">Create Post</a>
                          <a  href="/logout" class="btn my-2">Logout</a>
                          </span>
                          <a v-else href="/#/authorization" class="btn btn-primary my-2">Authorization</a>
                      </p>
                  </div>
              </div>
          </section>
        <div class="album py-5 bg-light">
            <div class="container">
                <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
                    <div class="card" id="app"  v-for="post in posts.Posts">
                        <div class="card-body">
                            <p class="card-text">{{ post.Title }}</p>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a v-bind:href="'/#/post/'+ post.Id"  class="btn btn-sm btn-outline-secondary">View</a>
                                    <a v-bind:href="'/#/edit/'+ post.Id" class="btn btn-sm btn-outline-secondary">Edit</a>
                                    <button v-on:click="deletePost(post.Id)" class="btn btn-sm btn-outline-secondary">Delete</button>
                                </div>
                                <small class="text-muted">{{ post.UserId }}</small>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
  </div>
</template>

<script>
export default {
  name: 'Index',
  data () {

    return {
    posts: [],
    }
  },
   mounted() {
      fetch("http://localhost:8000/posts")
        .then(response =>response.json())
        .then((data)=>{
         this.posts = data;
        })
    },
      watch: {
        '$route' (to,from){
          this.deletePost()
        }
      },

   methods: {
        deletePost(id) {
       if (confirm ("Delete post?")) {
            fetch("http://localhost:8000/deletePost/"+id, {
                method:"DELETE"
            })
            event.preventDefault()

            window.location = '/'}
        },
         mess: function (message, event) {
             if (event) {
                  event.preventDefault()
                }
                  alert(message);

                 },

   }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->

<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
