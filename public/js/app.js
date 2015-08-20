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

    var aux = data.date.split("-");
    var year = aux[0];
    var month = aux[1];
    var day = aux[2];

    $( "#day" ).text( day );
    $( "#month" ).text( month );
    $( "#year" ).text( year );
    $( "#description" ).text( data.description );

    $("#show-form").show();
  }).fail(function() {
    $("#create-form").show();
    $("#show-form").hide();
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
