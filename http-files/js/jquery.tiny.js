var Zepto = function() {
    function t(t) {
        return null == t ? String(t) : Y[X.call(t)] || "object"
    }

    function e(e) {
        return "function" == t(e)
    }

    function i(t) {
        return null != t && t == t.window
    }

    function n(t) {
        return null != t && t.nodeType == t.DOCUMENT_NODE
    }

    function r(e) {
        return "object" == t(e)
    }

    function o(t) {
        return r(t) && !i(t) && t.__proto__ == Object.prototype
    }

    function a(t) {
        return t instanceof Array
    }

    function s(t) {
        return "number" == typeof t.length
    }

    function u(t) {
        return M.call(t, function(t) {
            return null != t
        })
    }

    function c(t) {
        return t.length > 0 ? N.fn.concat.apply([], t) : t
    }

    function l(t) {
        return t.replace(/::/g, "/").replace(/([A-Z]+)([A-Z][a-z])/g, "$1_$2").replace(/([a-z\d])([A-Z])/g, "$1_$2").replace(/_/g, "-").toLowerCase()
    }

    function d(t) {
        return t in R ? R[t] : R[t] = new RegExp("(^|\\s)" + t + "(\\s|$)")
    }

    function f(t, e) {
        return "number" != typeof e || j[l(t)] ? e : e + "px"
    }

    function h(t) {
        var e, i;
        return B[t] || (e = D.createElement(t),
                D.body.appendChild(e),
                i = I(e, "").getPropertyValue("display"),
                e.parentNode.removeChild(e),
                "none" == i && (i = "block"),
                B[t] = i),
            B[t]
    }

    function p(t) {
        return "children" in t ? _.call(t.children) : N.map(t.childNodes, function(t) {
            return 1 == t.nodeType ? t : void 0
        })
    }

    function g(t, e, i) {
        for (S in e)
            i && (o(e[S]) || a(e[S])) ? (o(e[S]) && !o(t[S]) && (t[S] = {}),
                a(e[S]) && !a(t[S]) && (t[S] = []),
                g(t[S], e[S], i)) : e[S] !== C && (t[S] = e[S])
    }

    function m(t, e) {
        return e === C ? N(t) : N(t).filter(e)
    }

    function b(t, i, n, r) {
        return e(i) ? i.call(t, n, r) : i
    }

    function v(t, e, i) {
        null == i ? t.removeAttribute(e) : t.setAttribute(e, i)
    }

    function y(t, e) {
        var i = t.className,
            n = i && i.baseVal !== C;
        return e === C ? n ? i.baseVal : i : void(n ? i.baseVal = e : t.className = e)
    }

    function w(t) {
        var e;
        try {
            return t ? "true" == t || ("false" == t ? !1 : "null" == t ? null : isNaN(e = Number(t)) ? /^[\[\{]/.test(t) ? N.parseJSON(t) : t : e) : t
        } catch (i) {
            return t
        }
    }

    function x(t, e) {
        e(t);
        for (var i in t.childNodes)
            x(t.childNodes[i], e)
    }
    var C, S, N, T, E, k, A = [],
        _ = A.slice,
        M = A.filter,
        D = window.document,
        B = {},
        R = {},
        I = D.defaultView.getComputedStyle,
        j = {
            "column-count": 1,
            columns: 1,
            "font-weight": 1,
            "line-height": 1,
            opacity: 1,
            "z-index": 1,
            zoom: 1
        },
        P = /^\s*<(\w+|!)[^>]*>/,
        O = /<(?!area|br|col|embed|hr|img|input|link|meta|param)(([\w:]+)[^>]*)\/>/gi,
        H = /^(?:body|html)$/i,
        $ = ["val", "css", "html", "text", "data", "width", "height", "offset"],
        L = ["after", "prepend", "before", "append"],
        Z = D.createElement("table"),
        q = D.createElement("tr"),
        F = {
            tr: D.createElement("tbody"),
            tbody: Z,
            thead: Z,
            tfoot: Z,
            td: q,
            th: q,
            "*": D.createElement("div")
        },
        V = /complete|loaded|interactive/,
        z = /^\.([\w-]+)$/,
        U = /^#([\w-]*)$/,
        J = /^[\w-]+$/,
        Y = {},
        X = Y.toString,
        W = {},
        Q = D.createElement("div");
    return W.matches = function(t, e) {
            if (!t || 1 !== t.nodeType)
                return !1;
            var i = t.webkitMatchesSelector || t.mozMatchesSelector || t.oMatchesSelector || t.matchesSelector;
            if (i)
                return i.call(t, e);
            var n, r = t.parentNode,
                o = !r;
            return o && (r = Q).appendChild(t),
                n = ~W.qsa(r, e).indexOf(t),
                o && Q.removeChild(t),
                n
        },
        E = function(t) {
            return t.replace(/-+(.)?/g, function(t, e) {
                return e ? e.toUpperCase() : ""
            })
        },
        k = function(t) {
            return M.call(t, function(e, i) {
                return t.indexOf(e) == i
            })
        },
        W.fragment = function(t, e, i) {
            t.replace && (t = t.replace(O, "<$1></$2>")),
                e === C && (e = P.test(t) && RegExp.$1),
                e in F || (e = "*");
            var n, r, a = F[e];
            return a.innerHTML = "" + t,
                r = N.each(_.call(a.childNodes), function() {
                    a.removeChild(this)
                }),
                o(i) && (n = N(r),
                    N.each(i, function(t, e) {
                        $.indexOf(t) > -1 ? n[t](e) : n.attr(t, e)
                    })),
                r
        },
        W.Z = function(t, e) {
            return t = t || [],
                t.__proto__ = N.fn,
                t.selector = e || "",
                t
        },
        W.isZ = function(t) {
            return t instanceof W.Z
        },
        W.init = function(t, i) {
            if (t) {
                if (e(t))
                    return N(D).ready(t);
                if (W.isZ(t))
                    return t;
                var n;
                if (a(t))
                    n = u(t);
                else if (r(t))
                    n = [o(t) ? N.extend({}, t) : t],
                    t = null;
                else if (P.test(t))
                    n = W.fragment(t.trim(), RegExp.$1, i),
                    t = null;
                else {
                    if (i !== C)
                        return N(i).find(t);
                    n = W.qsa(D, t)
                }
                return W.Z(n, t)
            }
            return W.Z()
        },
        N = function(t, e) {
            return W.init(t, e)
        },
        N.extend = function(t) {
            var e, i = _.call(arguments, 1);
            return "boolean" == typeof t && (e = t,
                    t = i.shift()),
                i.forEach(function(i) {
                    g(t, i, e)
                }),
                t
        },
        W.qsa = function(t, e) {
            var i;
            return n(t) && U.test(e) ? (i = t.getElementById(RegExp.$1)) ? [i] : [] : 1 !== t.nodeType && 9 !== t.nodeType ? [] : _.call(z.test(e) ? t.getElementsByClassName(RegExp.$1) : J.test(e) ? t.getElementsByTagName(e) : t.querySelectorAll(e))
        },
        N.contains = function(t, e) {
            return t !== e && t.contains(e)
        },
        N.type = t,
        N.isFunction = e,
        N.isWindow = i,
        N.isArray = a,
        N.isPlainObject = o,
        N.isEmptyObject = function(t) {
            var e;
            for (e in t)
                return !1;
            return !0
        },
        N.inArray = function(t, e, i) {
            return A.indexOf.call(e, t, i)
        },
        N.camelCase = E,
        N.trim = function(t) {
            return t.trim()
        },
        N.uuid = 0,
        N.support = {},
        N.expr = {},
        N.map = function(t, e) {
            var i, n, r, o = [];
            if (s(t))
                for (n = 0; n < t.length; n++)
                    i = e(t[n], n),
                    null != i && o.push(i);
            else
                for (r in t)
                    i = e(t[r], r),
                    null != i && o.push(i);
            return c(o)
        },
        N.each = function(t, e) {
            var i, n;
            if (s(t)) {
                for (i = 0; i < t.length; i++)
                    if (e.call(t[i], i, t[i]) === !1)
                        return t
            } else
                for (n in t)
                    if (e.call(t[n], n, t[n]) === !1)
                        return t;
            return t
        },
        N.grep = function(t, e) {
            return M.call(t, e)
        },
        window.JSON && (N.parseJSON = JSON.parse),
        N.each("Boolean Number String Function Array Date RegExp Object Error".split(" "), function(t, e) {
            Y["[object " + e + "]"] = e.toLowerCase()
        }),
        N.fn = {
            forEach: A.forEach,
            reduce: A.reduce,
            push: A.push,
            sort: A.sort,
            indexOf: A.indexOf,
            concat: A.concat,
            map: function(t) {
                return N(N.map(this, function(e, i) {
                    return t.call(e, i, e)
                }))
            },
            slice: function() {
                return N(_.apply(this, arguments))
            },
            ready: function(t) {
                return V.test(D.readyState) ? t(N) : D.addEventListener("DOMContentLoaded", function() {
                        t(N)
                    }, !1),
                    this
            },
            get: function(t) {
                return t === C ? _.call(this) : this[t >= 0 ? t : t + this.length]
            },
            toArray: function() {
                return this.get()
            },
            size: function() {
                return this.length
            },
            remove: function() {
                return this.each(function() {
                    null != this.parentNode && this.parentNode.removeChild(this)
                })
            },
            each: function(t) {
                return A.every.call(this, function(e, i) {
                        return t.call(e, i, e) !== !1
                    }),
                    this
            },
            filter: function(t) {
                return e(t) ? this.not(this.not(t)) : N(M.call(this, function(e) {
                    return W.matches(e, t)
                }))
            },
            add: function(t, e) {
                return N(k(this.concat(N(t, e))))
            },
            is: function(t) {
                return this.length > 0 && W.matches(this[0], t)
            },
            not: function(t) {
                var i = [];
                if (e(t) && t.call !== C)
                    this.each(function(e) {
                        t.call(this, e) || i.push(this)
                    });
                else {
                    var n = "string" == typeof t ? this.filter(t) : s(t) && e(t.item) ? _.call(t) : N(t);
                    this.forEach(function(t) {
                        n.indexOf(t) < 0 && i.push(t)
                    })
                }
                return N(i)
            },
            has: function(t) {
                return this.filter(function() {
                    return r(t) ? N.contains(this, t) : N(this).find(t).size()
                })
            },
            eq: function(t) {
                return -1 === t ? this.slice(t) : this.slice(t, +t + 1)
            },
            first: function() {
                var t = this[0];
                return t && !r(t) ? t : N(t)
            },
            last: function() {
                var t = this[this.length - 1];
                return t && !r(t) ? t : N(t)
            },
            find: function(t) {
                var e, i = this;
                return e = "object" == typeof t ? N(t).filter(function() {
                    var t = this;
                    return A.some.call(i, function(e) {
                        return N.contains(e, t)
                    })
                }) : 1 == this.length ? N(W.qsa(this[0], t)) : this.map(function() {
                    return W.qsa(this, t)
                })
            },
            closest: function(t, e) {
                var i = this[0],
                    r = !1;
                for ("object" == typeof t && (r = N(t)); i && !(r ? r.indexOf(i) >= 0 : W.matches(i, t));)
                    i = i !== e && !n(i) && i.parentNode;
                return N(i)
            },
            parents: function(t) {
                for (var e = [], i = this; i.length > 0;)
                    i = N.map(i, function(t) {
                        return (t = t.parentNode) && !n(t) && e.indexOf(t) < 0 ? (e.push(t),
                            t) : void 0
                    });
                return m(e, t)
            },
            parent: function(t) {
                return m(k(this.pluck("parentNode")), t)
            },
            children: function(t) {
                return m(this.map(function() {
                    return p(this)
                }), t)
            },
            contents: function() {
                return this.map(function() {
                    return _.call(this.childNodes)
                })
            },
            siblings: function(t) {
                return m(this.map(function(t, e) {
                    return M.call(p(e.parentNode), function(t) {
                        return t !== e
                    })
                }), t)
            },
            empty: function() {
                return this.each(function() {
                    this.innerHTML = ""
                })
            },
            pluck: function(t) {
                return N.map(this, function(e) {
                    return e[t]
                })
            },
            show: function() {
                return this.each(function() {
                    "none" == this.style.display && (this.style.display = null),
                        "none" == I(this, "").getPropertyValue("display") && (this.style.display = h(this.nodeName))
                })
            },
            replaceWith: function(t) {
                return this.before(t).remove()
            },
            wrap: function(t) {
                var i = e(t);
                if (this[0] && !i)
                    var n = N(t).get(0),
                        r = n.parentNode || this.length > 1;
                return this.each(function(e) {
                    N(this).wrapAll(i ? t.call(this, e) : r ? n.cloneNode(!0) : n)
                })
            },
            wrapAll: function(t) {
                if (this[0]) {
                    N(this[0]).before(t = N(t));
                    for (var e;
                        (e = t.children()).length;)
                        t = e.first();
                    N(t).append(this)
                }
                return this
            },
            wrapInner: function(t) {
                var i = e(t);
                return this.each(function(e) {
                    var n = N(this),
                        r = n.contents(),
                        o = i ? t.call(this, e) : t;
                    r.length ? r.wrapAll(o) : n.append(o)
                })
            },
            unwrap: function() {
                return this.parent().each(function() {
                        N(this).replaceWith(N(this).children())
                    }),
                    this
            },
            clone: function() {
                return this.map(function() {
                    return this.cloneNode(!0)
                })
            },
            hide: function() {
                return this.css("display", "none")
            },
            toggle: function(t) {
                return this.each(function() {
                    var e = N(this);
                    (t === C ? "none" == e.css("display") : t) ? e.show(): e.hide()
                })
            },
            prev: function(t) {
                return N(this.pluck("previousElementSibling")).filter(t || "*")
            },
            next: function(t) {
                return N(this.pluck("nextElementSibling")).filter(t || "*")
            },
            html: function(t) {
                return t === C ? this.length > 0 ? this[0].innerHTML : null : this.each(function(e) {
                    var i = this.innerHTML;
                    N(this).empty().append(b(this, t, e, i))
                })
            },
            text: function(t) {
                return t === C ? this.length > 0 ? this[0].textContent : null : this.each(function() {
                    this.textContent = t
                })
            },
            attr: function(t, e) {
                var i;
                return "string" == typeof t && e === C ? 0 == this.length || 1 !== this[0].nodeType ? C : "value" == t && "INPUT" == this[0].nodeName ? this.val() : !(i = this[0].getAttribute(t)) && t in this[0] ? this[0][t] : i : this.each(function(i) {
                    if (1 === this.nodeType)
                        if (r(t))
                            for (S in t)
                                v(this, S, t[S]);
                        else
                            v(this, t, b(this, e, i, this.getAttribute(t)))
                })
            },
            removeAttr: function(t) {
                return this.each(function() {
                    1 === this.nodeType && v(this, t)
                })
            },
            prop: function(t, e) {
                return e === C ? this[0] && this[0][t] : this.each(function(i) {
                    this[t] = b(this, e, i, this[t])
                })
            },
            data: function(t, e) {
                var i = this.attr("data-" + l(t), e);
                return null !== i ? w(i) : C
            },
            val: function(t) {
                return t === C ? this[0] && (this[0].multiple ? N(this[0]).find("option").filter(function() {
                    return this.selected
                }).pluck("value") : this[0].value) : this.each(function(e) {
                    this.value = b(this, t, e, this.value)
                })
            },
            offset: function(t) {
                if (t)
                    return this.each(function(e) {
                        var i = N(this),
                            n = b(this, t, e, i.offset()),
                            r = i.offsetParent().offset(),
                            o = {
                                top: n.top - r.top,
                                left: n.left - r.left
                            };
                        "static" == i.css("position") && (o.position = "relative"),
                            i.css(o)
                    });
                if (0 == this.length)
                    return null;
                var e = this[0].getBoundingClientRect();
                return {
                    left: e.left + window.pageXOffset,
                    top: e.top + window.pageYOffset,
                    width: Math.round(e.width),
                    height: Math.round(e.height)
                }
            },
            css: function(e, i) {
                if (arguments.length < 2 && "string" == typeof e)
                    return this[0] && (this[0].style[E(e)] || I(this[0], "").getPropertyValue(e));
                var n = "";
                if ("string" == t(e))
                    i || 0 === i ? n = l(e) + ":" + f(e, i) : this.each(function() {
                        this.style.removeProperty(l(e))
                    });
                else
                    for (S in e)
                        e[S] || 0 === e[S] ? n += l(S) + ":" + f(S, e[S]) + ";" : this.each(function() {
                            this.style.removeProperty(l(S))
                        });
                return this.each(function() {
                    this.style.cssText += ";" + n
                })
            },
            index: function(t) {
                return t ? this.indexOf(N(t)[0]) : this.parent().children().indexOf(this[0])
            },
            hasClass: function(t) {
                return A.some.call(this, function(t) {
                    return this.test(y(t))
                }, d(t))
            },
            addClass: function(t) {
                return this.each(function(e) {
                    T = [];
                    var i = y(this),
                        n = b(this, t, e, i);
                    n.split(/\s+/g).forEach(function(t) {
                            N(this).hasClass(t) || T.push(t)
                        }, this),
                        T.length && y(this, i + (i ? " " : "") + T.join(" "))
                })
            },
            removeClass: function(t) {
                return this.each(function(e) {
                    return t === C ? y(this, "") : (T = y(this),
                        b(this, t, e, T).split(/\s+/g).forEach(function(t) {
                            T = T.replace(d(t), " ")
                        }),
                        void y(this, T.trim()))
                })
            },
            toggleClass: function(t, e) {
                return this.each(function(i) {
                    var n = N(this),
                        r = b(this, t, i, y(this));
                    r.split(/\s+/g).forEach(function(t) {
                        (e === C ? !n.hasClass(t) : e) ? n.addClass(t): n.removeClass(t)
                    })
                })
            },
            scrollTop: function() {
                return this.length ? "scrollTop" in this[0] ? this[0].scrollTop : this[0].scrollY : void 0
            },
            position: function() {
                if (this.length) {
                    var t = this[0],
                        e = this.offsetParent(),
                        i = this.offset(),
                        n = H.test(e[0].nodeName) ? {
                            top: 0,
                            left: 0
                        } : e.offset();
                    return i.top -= parseFloat(N(t).css("margin-top")) || 0,
                        i.left -= parseFloat(N(t).css("margin-left")) || 0,
                        n.top += parseFloat(N(e[0]).css("border-top-width")) || 0,
                        n.left += parseFloat(N(e[0]).css("border-left-width")) || 0, {
                            top: i.top - n.top,
                            left: i.left - n.left
                        }
                }
            },
            offsetParent: function() {
                return this.map(function() {
                    for (var t = this.offsetParent || D.body; t && !H.test(t.nodeName) && "static" == N(t).css("position");)
                        t = t.offsetParent;
                    return t
                })
            }
        },
        N.fn.detach = N.fn.remove, ["width", "height"].forEach(function(t) {
            N.fn[t] = function(e) {
                var r, o = this[0],
                    a = t.replace(/./, function(t) {
                        return t[0].toUpperCase()
                    });
                return e === C ? i(o) ? o["inner" + a] : n(o) ? o.documentElement["offset" + a] : (r = this.offset()) && r[t] : this.each(function(i) {
                    o = N(this),
                        o.css(t, b(this, e, i, o[t]()))
                })
            }
        }),
        L.forEach(function(e, i) {
            var n = i % 2;
            N.fn[e] = function() {
                    var e, r, o = N.map(arguments, function(i) {
                            return e = t(i),
                                "object" == e || "array" == e || null == i ? i : W.fragment(i)
                        }),
                        a = this.length > 1;
                    return o.length < 1 ? this : this.each(function(t, e) {
                        r = n ? e : e.parentNode,
                            e = 0 == i ? e.nextSibling : 1 == i ? e.firstChild : 2 == i ? e : null,
                            o.forEach(function(t) {
                                if (a)
                                    t = t.cloneNode(!0);
                                else if (!r)
                                    return N(t).remove();
                                x(r.insertBefore(t, e), function(t) {
                                    null == t.nodeName || "SCRIPT" !== t.nodeName.toUpperCase() || t.type && "text/javascript" !== t.type || t.src || window.eval.call(window, t.innerHTML)
                                })
                            })
                    })
                },
                N.fn[n ? e + "To" : "insert" + (i ? "Before" : "After")] = function(t) {
                    return N(t)[e](this),
                        this
                }
        }),
        W.Z.prototype = N.fn,
        W.uniq = k,
        W.deserializeValue = w,
        N.zepto = W,
        N
}();
window.Zepto = Zepto,
    "$" in window || (window.$ = Zepto),
    function(t) {
        function e(t) {
            var e = this.os = {},
                i = this.browser = {},
                n = t.match(/WebKit\/([\d.]+)/),
                r = t.match(/(Android)\s+([\d.]+)/),
                o = t.match(/(iPad).*OS\s([\d_]+)/),
                a = !o && t.match(/(iPhone\sOS)\s([\d_]+)/),
                s = t.match(/(webOS|hpwOS)[\s\/]([\d.]+)/),
                u = s && t.match(/TouchPad/),
                c = t.match(/Kindle\/([\d.]+)/),
                l = t.match(/Silk\/([\d._]+)/),
                d = t.match(/(BlackBerry).*Version\/([\d.]+)/),
                f = t.match(/(BB10).*Version\/([\d.]+)/),
                h = t.match(/(RIM\sTablet\sOS)\s([\d.]+)/),
                p = t.match(/PlayBook/),
                g = t.match(/Chrome\/([\d.]+)/) || t.match(/CriOS\/([\d.]+)/),
                m = t.match(/Firefox\/([\d.]+)/);
            (i.webkit = !!n) && (i.version = n[1]),
            r && (e.android = !0,
                    e.version = r[2]),
                a && (e.ios = e.iphone = !0,
                    e.version = a[2].replace(/_/g, ".")),
                o && (e.ios = e.ipad = !0,
                    e.version = o[2].replace(/_/g, ".")),
                s && (e.webos = !0,
                    e.version = s[2]),
                u && (e.touchpad = !0),
                d && (e.blackberry = !0,
                    e.version = d[2]),
                f && (e.bb10 = !0,
                    e.version = f[2]),
                h && (e.rimtabletos = !0,
                    e.version = h[2]),
                p && (i.playbook = !0),
                c && (e.kindle = !0,
                    e.version = c[1]),
                l && (i.silk = !0,
                    i.version = l[1]), !l && e.android && t.match(/Kindle Fire/) && (i.silk = !0),
                g && (i.chrome = !0,
                    i.version = g[1]),
                m && (i.firefox = !0,
                    i.version = m[1]),
                e.tablet = !!(o || p || r && !t.match(/Mobile/) || m && t.match(/Tablet/)),
                e.phone = !(e.tablet || !(r || a || s || d || f || g && t.match(/Android/) || g && t.match(/CriOS\/([\d.]+)/) || m && t.match(/Mobile/)))
        }
        e.call(t, navigator.userAgent),
            t.__detect = e
    }(Zepto),
    function(t) {
        function e(t) {
            return t._zid || (t._zid = h++)
        }

        function i(t, i, o, a) {
            if (i = n(i),
                i.ns)
                var s = r(i.ns);
            return (f[e(t)] || []).filter(function(t) {
                return !(!t || i.e && t.e != i.e || i.ns && !s.test(t.ns) || o && e(t.fn) !== e(o) || a && t.sel != a)
            })
        }

        function n(t) {
            var e = ("" + t).split(".");
            return {
                e: e[0],
                ns: e.slice(1).sort().join(" ")
            }
        }

        function r(t) {
            return new RegExp("(?:^| )" + t.replace(" ", " .* ?") + "(?: |$)")
        }

        function o(e, i, n) {
            "string" != t.type(e) ? t.each(e, n) : e.split(/\s/).forEach(function(t) {
                n(t, i)
            })
        }

        function a(t, e) {
            return t.del && ("focus" == t.e || "blur" == t.e) || !!e
        }

        function s(t) {
            return g[t] || t
        }

        function u(i, r, u, c, l, d) {
            var h = e(i),
                p = f[h] || (f[h] = []);
            o(r, u, function(e, r) {
                var o = n(e);
                o.fn = r,
                    o.sel = c,
                    o.e in g && (r = function(e) {
                        var i = e.relatedTarget;
                        return !i || i !== this && !t.contains(this, i) ? o.fn.apply(this, arguments) : void 0
                    }),
                    o.del = l && l(r, e);
                var u = o.del || r;
                o.proxy = function(t) {
                        var e = u.apply(i, [t].concat(t.data));
                        return e === !1 && (t.preventDefault(),
                                t.stopPropagation()),
                            e
                    },
                    o.i = p.length,
                    p.push(o),
                    i.addEventListener(s(o.e), o.proxy, a(o, d))
            })
        }

        function c(t, n, r, u, c) {
            var l = e(t);
            o(n || "", r, function(e, n) {
                i(t, e, n, u).forEach(function(e) {
                    delete f[l][e.i],
                        t.removeEventListener(s(e.e), e.proxy, a(e, c))
                })
            })
        }

        function l(e) {
            var i, n = {
                originalEvent: e
            };
            for (i in e)
                v.test(i) || void 0 === e[i] || (n[i] = e[i]);
            return t.each(y, function(t, i) {
                    n[t] = function() {
                            return this[i] = m,
                                e[t].apply(e, arguments)
                        },
                        n[i] = b
                }),
                n
        }

        function d(t) {
            if (!("defaultPrevented" in t)) {
                t.defaultPrevented = !1;
                var e = t.preventDefault;
                t.preventDefault = function() {
                    this.defaultPrevented = !0,
                        e.call(this)
                }
            }
        }
        var f = (t.zepto.qsa, {}),
            h = 1,
            p = {},
            g = {
                mouseenter: "mouseover",
                mouseleave: "mouseout"
            };
        p.click = p.mousedown = p.mouseup = p.mousemove = "MouseEvents",
            t.event = {
                add: u,
                remove: c
            },
            t.proxy = function(i, n) {
                if (t.isFunction(i)) {
                    var r = function() {
                        return i.apply(n, arguments)
                    };
                    return r._zid = e(i),
                        r
                }
                if ("string" == typeof n)
                    return t.proxy(i[n], i);
                throw new TypeError("expected function")
            },
            t.fn.bind = function(t, e) {
                return this.each(function() {
                    u(this, t, e)
                })
            },
            t.fn.unbind = function(t, e) {
                return this.each(function() {
                    c(this, t, e)
                })
            },
            t.fn.one = function(t, e) {
                return this.each(function(i, n) {
                    u(this, t, e, null, function(t, e) {
                        return function() {
                            var i = t.apply(n, arguments);
                            return c(n, e, t),
                                i
                        }
                    })
                })
            };
        var m = function() {
                return !0
            },
            b = function() {
                return !1
            },
            v = /^([A-Z]|layer[XY]$)/,
            y = {
                preventDefault: "isDefaultPrevented",
                stopImmediatePropagation: "isImmediatePropagationStopped",
                stopPropagation: "isPropagationStopped"
            };
        t.fn.delegate = function(e, i, n) {
                return this.each(function(r, o) {
                    u(o, i, n, e, function(i) {
                        return function(n) {
                            var r, a = t(n.target).closest(e, o).get(0);
                            return a ? (r = t.extend(l(n), {
                                    currentTarget: a,
                                    liveFired: o
                                }),
                                i.apply(a, [r].concat([].slice.call(arguments, 1)))) : void 0
                        }
                    })
                })
            },
            t.fn.undelegate = function(t, e, i) {
                return this.each(function() {
                    c(this, e, i, t)
                })
            },
            t.fn.live = function(e, i) {
                return t(document.body).delegate(this.selector, e, i),
                    this
            },
            t.fn.die = function(e, i) {
                return t(document.body).undelegate(this.selector, e, i),
                    this
            },
            t.fn.on = function(e, i, n) {
                return !i || t.isFunction(i) ? this.bind(e, i || n) : this.delegate(i, e, n)
            },
            t.fn.off = function(e, i, n) {
                return !i || t.isFunction(i) ? this.unbind(e, i || n) : this.undelegate(i, e, n)
            },
            t.fn.trigger = function(e, i) {
                return ("string" == typeof e || t.isPlainObject(e)) && (e = t.Event(e)),
                    d(e),
                    e.data = i,
                    this.each(function() {
                        "dispatchEvent" in this && this.dispatchEvent(e)
                    })
            },
            t.fn.triggerHandler = function(e, n) {
                var r, o;
                return this.each(function(a, s) {
                        r = l("string" == typeof e ? t.Event(e) : e),
                            r.data = n,
                            r.target = s,
                            t.each(i(s, e.type || e), function(t, e) {
                                return o = e.proxy(r),
                                    r.isImmediatePropagationStopped() ? !1 : void 0
                            })
                    }),
                    o
            },
            "focusin focusout load resize scroll unload click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select keydown keypress keyup error".split(" ").forEach(function(e) {
                t.fn[e] = function(t) {
                    return t ? this.bind(e, t) : this.trigger(e)
                }
            }), ["focus", "blur"].forEach(function(e) {
                t.fn[e] = function(t) {
                    return t ? this.bind(e, t) : this.each(function() {
                            try {
                                this[e]()
                            } catch (t) {}
                        }),
                        this
                }
            }),
            t.Event = function(t, e) {
                "string" != typeof t && (e = t,
                    t = e.type);
                var i = document.createEvent(p[t] || "Events"),
                    n = !0;
                if (e)
                    for (var r in e)
                        "bubbles" == r ? n = !!e[r] : i[r] = e[r];
                return i.initEvent(t, n, !0, null, null, null, null, null, null, null, null, null, null, null, null),
                    i.isDefaultPrevented = function() {
                        return this.defaultPrevented
                    },
                    i
            }
    }(Zepto),
    function(t) {
        function e(e, i, n) {
            var r = t.Event(i);
            return t(e).trigger(r, n), !r.defaultPrevented
        }

        function i(t, i, n, r) {
            return t.global ? e(i || v, n, r) : void 0
        }

        function n(e) {
            e.global && 0 === t.active++ && i(e, null, "ajaxStart")
        }

        function r(e) {
            e.global && !--t.active && i(e, null, "ajaxStop")
        }

        function o(t, e) {
            var n = e.context;
            return e.beforeSend.call(n, t, e) === !1 || i(e, n, "ajaxBeforeSend", [t, e]) === !1 ? !1 : void i(e, n, "ajaxSend", [t, e])
        }

        function a(t, e, n) {
            var r = n.context,
                o = "success";
            n.success.call(r, t, o, e),
                i(n, r, "ajaxSuccess", [e, n, t]),
                u(o, e, n)
        }

        function s(t, e, n, r) {
            var o = r.context;
            r.error.call(o, n, e, t),
                i(r, o, "ajaxError", [n, r, t]),
                u(e, n, r)
        }

        function u(t, e, n) {
            var o = n.context;
            n.complete.call(o, e, t),
                i(n, o, "ajaxComplete", [e, n]),
                r(n)
        }

        function c() {}

        function l(t) {
            return t && (t = t.split(";", 2)[0]),
                t && (t == S ? "html" : t == C ? "json" : w.test(t) ? "script" : x.test(t) && "xml") || "text"
        }

        function d(t, e) {
            return (t + "&" + e).replace(/[&?]{1,2}/, "?")
        }

        function f(e) {
            e.processData && e.data && "string" != t.type(e.data) && (e.data = t.param(e.data, e.traditional)), !e.data || e.type && "GET" != e.type.toUpperCase() || (e.url = d(e.url, e.data))
        }

        function h(e, i, n, r) {
            var o = !t.isFunction(i);
            return {
                url: e,
                data: o ? i : void 0,
                success: o ? t.isFunction(n) ? n : void 0 : i,
                dataType: o ? r || n : n
            }
        }

        function p(e, i, n, r) {
            var o, a = t.isArray(i);
            t.each(i, function(i, s) {
                o = t.type(s),
                    r && (i = n ? r : r + "[" + (a ? "" : i) + "]"), !r && a ? e.add(s.name, s.value) : "array" == o || !n && "object" == o ? p(e, s, n, i) : e.add(i, s)
            })
        }
        var g, m, b = 0,
            v = window.document,
            y = /<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi,
            w = /^(?:text|application)\/javascript/i,
            x = /^(?:text|application)\/xml/i,
            C = "application/json",
            S = "text/html",
            N = /^\s*$/;
        t.active = 0,
            t.ajaxJSONP = function(e) {
                if (!("type" in e))
                    return t.ajax(e);
                var i, n = "jsonp" + ++b,
                    r = v.createElement("script"),
                    u = function() {
                        clearTimeout(i),
                            t(r).remove(),
                            delete window[n]
                    },
                    l = function(t) {
                        u(),
                            t && "timeout" != t || (window[n] = c),
                            s(null, t || "abort", d, e)
                    },
                    d = {
                        abort: l
                    };
                return o(d, e) === !1 ? (l("abort"), !1) : (window[n] = function(t) {
                        u(),
                            a(t, d, e)
                    },
                    r.onerror = function() {
                        l("error")
                    },
                    r.src = e.url.replace(/=\?/, "=" + n),
                    t("head").append(r),
                    e.timeout > 0 && (i = setTimeout(function() {
                        l("timeout")
                    }, e.timeout)),
                    d)
            },
            t.ajaxSettings = {
                type: "GET",
                beforeSend: c,
                success: c,
                error: c,
                complete: c,
                context: null,
                global: !0,
                xhr: function() {
                    return new window.XMLHttpRequest
                },
                accepts: {
                    script: "text/javascript, application/javascript",
                    json: C,
                    xml: "application/xml, text/xml",
                    html: S,
                    text: "text/plain"
                },
                crossDomain: !1,
                timeout: 0,
                processData: !0,
                cache: !0
            },
            t.ajax = function(e) {
                var i = t.extend({}, e || {});
                for (g in t.ajaxSettings)
                    void 0 === i[g] && (i[g] = t.ajaxSettings[g]);
                n(i),
                    urlProtcol = null,
                    urlDomain = null,
                    crosstest = /^([\w-]+:)?\/\/([^\/]+)/.test(i.url),
                    urlProtcol = RegExp.$1,
                    urlDomain = RegExp.$2;
                var r = "wappass.baidu.com";
                "http:" == urlProtcol && urlDomain == r && (i.url = i.url.replace(/^http:/, "https:")),
                    i.crossDomain || (i.crossDomain = crosstest && urlDomain != window.location.host),
                    i.url || (i.url = window.location.toString()),
                    f(i),
                    i.cache === !1 && (i.url = d(i.url, "_=" + Date.now()));
                var u = i.dataType,
                    h = /=\?/.test(i.url);
                if ("jsonp" == u || h)
                    return h || (i.url = d(i.url, "callback=?")),
                        t.ajaxJSONP(i);
                var p, b = i.accepts[u],
                    v = {},
                    y = /^([\w-]+:)\/\//.test(i.url) ? RegExp.$1 : window.location.protocol,
                    w = i.xhr();
                urlDomain == r && "http:" == window.location.protocol && (i.crossDomain = !0),
                    i.crossDomain || (v["X-Requested-With"] = "XMLHttpRequest"),
                    b && (v.Accept = b,
                        b.indexOf(",") > -1 && (b = b.split(",", 2)[0]),
                        w.overrideMimeType && w.overrideMimeType(b)),
                    (i.contentType || i.contentType !== !1 && i.data && "GET" != i.type.toUpperCase()) && (v["Content-Type"] = i.contentType || "application/x-www-form-urlencoded"),
                    i.headers = t.extend(v, i.headers || {}),
                    w.onreadystatechange = function() {
                        if (4 == w.readyState) {
                            w.onreadystatechange = c,
                                clearTimeout(p);
                            var e, n = !1;
                            if (w.status >= 200 && w.status < 300 || 304 == w.status || 0 == w.status && "file:" == y) {
                                u = u || l(w.getResponseHeader("content-type")),
                                    e = w.responseText;
                                try {
                                    "script" == u ? (1,
                                        eval)(e) : "xml" == u ? e = w.responseXML : "json" == u && (e = N.test(e) ? null : t.parseJSON(e))
                                } catch (r) {
                                    n = r
                                }
                                n ? s(n, "parsererror", w, i) : a(e, w, i)
                            } else
                                s(null, w.status ? "error" : "abort", w, i)
                        }
                    };
                var x = "async" in i ? i.async : !0;
                w.open(i.type, i.url, x),
                    urlDomain == r && (w.withCredentials = !0);
                for (m in i.headers)
                    w.setRequestHeader(m, i.headers[m]);
                return o(w, i) === !1 ? (w.abort(), !1) : (i.timeout > 0 && (p = setTimeout(function() {
                        w.onreadystatechange = c,
                            w.abort(),
                            s(null, "timeout", w, i)
                    }, i.timeout)),
                    w.send(i.data ? i.data : null),
                    w)
            },
            t.get = function() {
                return t.ajax(h.apply(null, arguments))
            },
            t.post = function() {
                var e = h.apply(null, arguments);
                return e.type = "POST",
                    t.ajax(e)
            },
            t.getJSON = function() {
                var e = h.apply(null, arguments);
                return e.dataType = "json",
                    t.ajax(e)
            },
            t.fn.load = function(e, i, n) {
                if (!this.length)
                    return this;
                var r, o = this,
                    a = e.split(/\s/),
                    s = h(e, i, n),
                    u = s.success;
                return a.length > 1 && (s.url = a[0],
                        r = a[1]),
                    s.success = function(e) {
                        o.html(r ? t("<div>").html(e.replace(y, "")).find(r) : e),
                            u && u.apply(o, arguments)
                    },
                    t.ajax(s),
                    this
            };
        var T = encodeURIComponent;
        t.param = function(t, e) {
            var i = [];
            return i.add = function(t, e) {
                    this.push(T(t) + "=" + T(e))
                },
                p(i, t, e),
                i.join("&").replace(/%20/g, "+")
        }
    }(Zepto),
    function(t) {
        t.fn.serializeArray = function() {
                var e, i = [];
                return t(Array.prototype.slice.call(this.get(0).elements)).each(function() {
                        e = t(this);
                        var n = e.attr("type");
                        "fieldset" != this.nodeName.toLowerCase() && !this.disabled && "submit" != n && "reset" != n && "button" != n && ("radio" != n && "checkbox" != n || this.checked) && i.push({
                            name: e.attr("name"),
                            value: e.val()
                        })
                    }),
                    i
            },
            t.fn.serialize = function() {
                var t = [];
                return this.serializeArray().forEach(function(e) {
                        t.push(encodeURIComponent(e.name) + "=" + encodeURIComponent(e.value))
                    }),
                    t.join("&")
            },
            t.fn.submit = function(e) {
                if (e)
                    this.bind("submit", e);
                else if (this.length) {
                    var i = t.Event("submit");
                    this.eq(0).trigger(i),
                        i.defaultPrevented || this.get(0).submit()
                }
                return this
            }
    }(Zepto);