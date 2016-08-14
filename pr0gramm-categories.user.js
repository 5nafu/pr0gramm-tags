// ==UserScript==
// @name         pr0gramm-categories
// @namespace    http://github.net/mopsalarm/pr0gramm-tags
// @version      0.1.0
// @description  Allow better searches
// @author       Mopsalarm
// @match        http://pr0gramm.com/*
// @match        https://pr0gramm.com/*
// @grant        none
// @run-at       document-idle
// ==/UserScript==

(function () {
    'use strict';

    var orgGet = p.api.get;
    p.api.get = function (endpoint, opts, success, error) {
        if (endpoint === "items.get") {
            var hasSpecialPrefix = (opts.tags || "")[0] === "?";
            if (hasSpecialPrefix) {
                opts.tags = opts.tags.slice(1);
                return jQuery.ajax({
                    type: "GET",
                    url: "//app.pr0gramm.com/api/categories/v1/general",
                    success: success,
                    error: error,
                    dataType: "json",
                    data: opts
                });
            }
        }

        return orgGet.apply(this, arguments);
    };
})();
