/// <reference path="jquery.d.ts" />
$(document).ready(function () {
    $.get("api/miningInfo", function (data) {
        $("#hashrate").html(data.networkhashps);
    });
});
