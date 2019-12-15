var app = new Vue(
    {
        el: "#app",
        data: {
            message: "test",
            cwd: "/",
            info: "",
            contents: []
        },
        created: function(){
            this.getContents()
            this.getInfo()
        },
        methods: {
            getInfo: function(){
                axios({
                    method: "get",
                    url: "/content",
                    params: {path: "docker-compose.yml"}
                })
                .then(function(respone){
                    this.app.info = respone.data.info
                })
                .catch(function(error){
                    console.log(error)
                })
            },
            getContents: function(){
                axios.get("/content")
                .then(function(respone){
                    console.log(respone)
                    this.app.contents = respone.data.contents
                })
                .catch(function(error){
                    console.log(error)
                })
                .finally(function(){
                    console.log("aa")
                })
            }
        }
    }
)