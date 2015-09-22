$( document ).ready(function() {

  Commit.user = $("#user").val();
  retrieveCommit();

  $("#create-button").click(function(event) {
    event.preventDefault();
    description = $('input[name=description]').val();
    date        = $('input[name=date]').val() + "T05:00:00.000Z";
    form = { description: description, date: date };
    $.post("/commitment/" + Commit.user, form, function(data) {
      data.date = date;
      retrieveCommit();
    });
  });

  $("#complete-button").click(function(event) {
    event.preventDefault();
    form = { id: Commit.id, status: "completed" };
    $.put("/commitment/" + Commit.user, form, function(data) {
      retrieveCommit();
    });
  });

  $("#discard-button").click(function(event) {
    event.preventDefault();
    form = { id: Commit.id, status: "discarded" };
    $.put("/commitment/" + Commit.user, form, function(data) {
      retrieveCommit();
    });
  });

});

var Commit = {};

var retrieveCommit = function() {
  $.get("/commitment/" + Commit.user, function(data) {
    $("#create-form").hide();    
    Commit.id = data.id;

    var date = new Date(data.date);
    $( "#day" ).text( date.getDate() );
    $( "#month" ).text( ("0" + (date.getMonth() + 1)).slice(-2) );
    $( "#year" ).text( date.getFullYear() );
    $( "#description" ).text( data.description );
    $("#show-form").show();

  }).fail(function() {
    retrieveP();
    $("#show-form").hide();
    $("#create-form").show();
    $("#create-form")[0].reset();
    
  });
};

var retrieveP = function() {
   $.get("/commitment/" + Commit.user + "/percent", function(data) {
    $("#label").text("Solo has cumplido " + data.porciento + "% proyecto(s) de tu 100% (" + data.p + " de " + data.max +").");

    $("progress").attr('value', data.porciento );
    //$("progress").attr('max',  100);
  }).fail(function() {
    console.log("retrieveP ")
  });

};


$.put = function(url, data, callback, type){
  if ( $.isFunction(data) ){
    type = type || callback,
    callback = data,
    data = {}
  }

  return $.ajax({
    url: url,
    type: 'PUT',
    success: callback,
    data: data,
    contentType: type
  });
}
