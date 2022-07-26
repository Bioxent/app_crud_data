function removeFromDb(item){
    fetch(`/delete?item=${item}`, {method: "Delete"}).then(res =>{
        if (res.status == 200){
            window.location.pathname = "/"
        }
    })
 }
