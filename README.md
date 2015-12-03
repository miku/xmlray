README
======

xmlray is a little x-ray thing for xml.

![](http://etc.usf.edu/clipart/22900/22922/ray_22922_sm.gif)

Toy Linux binary: [xmlray](https://github.com/miku/xmlray/releases/download/v0.0.2/xmlray)

Usage
-----

    $ xmlray fixtures/sample.xml
    /a
    /a/b
    /a/b/c
    /a/b
    /a/b/c

Note that this is similar to [xml2](http://dan.egnor.name/xml2/ref):

    $ xml2 < fixtures/sample.xml
    /a/b/c
    /a/b
    /a/b/c

If the xml2 flat format is sufficient, we will switch to it - however,
some use cases might be difficult to support with a more verbose format,
that lists *all* nodes.

    $ xmlray fixtures/oai.xml
    /record
    /record/header
    /record/header/identifier
    /record/header/identifier/#
    /record/header/datestamp
    /record/header/datestamp/#
    /record/header/setSpec
    /record/header/setSpec/#
    /record/metadata
    /record/metadata/dc
    /record/metadata/dc/@oai_dc
    /record/metadata/dc/@dc
    /record/metadata/dc/@xsi
    /record/metadata/dc/@schemaLocation
    /record/metadata/dc/title
    /record/metadata/dc/title/#
    /record/metadata/dc/creator
    /record/metadata/dc/creator/#
    /record/metadata/dc/subject
    /record/metadata/dc/subject/#
    /record/metadata/dc/subject
    /record/metadata/dc/subject/#
    /record/metadata/dc/description
    /record/metadata/dc/description/#
    /record/metadata/dc/description
    /record/metadata/dc/description/#
    /record/metadata/dc/date
    /record/metadata/dc/date/#
    /record/metadata/dc/type
    /record/metadata/dc/type/#
    /record/metadata/dc/identifier
    /record/metadata/dc/identifier/#
    /record
    /record/header
    /record/header/identifier
    /record/header/identifier/#
    /record/header/datestamp
    /record/header/datestamp/#
    /record/header/setSpec
    /record/header/setSpec/#
    /record/metadata
    /record/metadata/dc
    /record/metadata/dc/@oai_dc
    /record/metadata/dc/@dc
    /record/metadata/dc/@xsi
    /record/metadata/dc/@schemaLocation
    /record/metadata/dc/title
    /record/metadata/dc/title/#
    /record/metadata/dc/creator
    /record/metadata/dc/creator/#
    /record/metadata/dc/subject
    /record/metadata/dc/subject/#
    /record/metadata/dc/description
    /record/metadata/dc/description/#
    /record/metadata/dc/description
    /record/metadata/dc/description/#
    /record/metadata/dc/date
    /record/metadata/dc/date/#
    /record/metadata/dc/type
    /record/metadata/dc/type/#
    /record/metadata/dc/identifier
    /record/metadata/dc/identifier/#
    /record
    /record/header
    /record/header/identifier
    /record/header/identifier/#
    /record/header/datestamp
    /record/header/datestamp/#
    /record/header/setSpec
    /record/header/setSpec/#
    /record/metadata
    /record/metadata/dc
    /record/metadata/dc/@oai_dc
    /record/metadata/dc/@dc
    /record/metadata/dc/@xsi
    /record/metadata/dc/@schemaLocation
    /record/metadata/dc/title
    /record/metadata/dc/title/#
    /record/metadata/dc/creator
    /record/metadata/dc/creator/#
    /record/metadata/dc/subject
    /record/metadata/dc/subject/#
    /record/metadata/dc/subject
    /record/metadata/dc/subject/#
    /record/metadata/dc/description
    /record/metadata/dc/description/#
    /record/metadata/dc/date
    /record/metadata/dc/date/#
    /record/metadata/dc/type
    /record/metadata/dc/type/#
    /record/metadata/dc/identifier
    /record/metadata/dc/identifier/#

Overview of element usage:

    $ xmlray fixtures/oai.xml | sort | uniq -c | sort -nr
      5 /record/metadata/dc/subject/#
      5 /record/metadata/dc/subject
      5 /record/metadata/dc/description/#
      5 /record/metadata/dc/description
      3 /record/metadata/dc/type/#
      3 /record/metadata/dc/type
      3 /record/metadata/dc/title/#
      3 /record/metadata/dc/title
      3 /record/metadata/dc/identifier/#
      3 /record/metadata/dc/identifier
      3 /record/metadata/dc/date/#
      3 /record/metadata/dc/date
      3 /record/metadata/dc/creator/#
      3 /record/metadata/dc/creator
      3 /record/metadata/dc/@xsi
      3 /record/metadata/dc/@schemaLocation
      3 /record/metadata/dc/@oai_dc
      3 /record/metadata/dc/@dc
      3 /record/metadata/dc
      3 /record/metadata
      3 /record/header/setSpec/#
      3 /record/header/setSpec
      3 /record/header/identifier/#
      3 /record/header/identifier
      3 /record/header/datestamp/#
      3 /record/header/datestamp
      3 /record/header
      3 /record

Preliminary schema visitor:

    $ xmlray -visitor schema -path /record <(zcat fixtures/xMetaDissPlus.xml.gz)
    /record
    /record/header
    /record/header/datestamp
    /record/header/datestamp/#
    /record/header/identifier
    /record/header/identifier/#
    /record/metadata
    /record/metadata/xMetaDiss
    /record/metadata/xMetaDiss/@aiiso
    /record/metadata/xMetaDiss/@cc
    /record/metadata/xMetaDiss/@dc
    /record/metadata/xMetaDiss/@dcmitype
    /record/metadata/xMetaDiss/@dcterms
    /record/metadata/xMetaDiss/@dctypes
    /record/metadata/xMetaDiss/@ddb
    /record/metadata/xMetaDiss/@dini
    /record/metadata/xMetaDiss/@doi
    /record/metadata/xMetaDiss/@fabio
    /record/metadata/xMetaDiss/@foaf
    /record/metadata/xMetaDiss/@hdl
    /record/metadata/xMetaDiss/@oai
    /record/metadata/xMetaDiss/@ore
    /record/metadata/xMetaDiss/@pc
    /record/metadata/xMetaDiss/@prism
    /record/metadata/xMetaDiss/@rdf
    /record/metadata/xMetaDiss/@schemaLocation
    /record/metadata/xMetaDiss/@skos
    /record/metadata/xMetaDiss/@thesis
    /record/metadata/xMetaDiss/@urn
    /record/metadata/xMetaDiss/@xMetaDiss
    /record/metadata/xMetaDiss/@xmlns
    /record/metadata/xMetaDiss/@xsi
    /record/metadata/xMetaDiss/alternative
    /record/metadata/xMetaDiss/alternative/#
    /record/metadata/xMetaDiss/alternative/@lang
    /record/metadata/xMetaDiss/alternative/@type
    /record/metadata/xMetaDiss/contact
    /record/metadata/xMetaDiss/contact/@contactID
    /record/metadata/xMetaDiss/contributor/person/academicTitle
    /record/metadata/xMetaDiss/contributor/person/academicTitle/#
    /record/metadata/xMetaDiss/contributor/person/name/foreName
    /record/metadata/xMetaDiss/contributor/person/name/foreName/#
    /record/metadata/xMetaDiss/contributor/person/name/surName
    /record/metadata/xMetaDiss/contributor/person/name/surName/#
    /record/metadata/xMetaDiss/created
    /record/metadata/xMetaDiss/created/#
    /record/metadata/xMetaDiss/creator/#
    /record/metadata/xMetaDiss/dateAccepted
    /record/metadata/xMetaDiss/dateAccepted/#
    /record/metadata/xMetaDiss/dateAccepted/@type
    /record/metadata/xMetaDiss/degree
    /record/metadata/xMetaDiss/degree/grantor
    /record/metadata/xMetaDiss/degree/grantor/@type
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/department
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/department/name
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/department/name/#
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/name
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/name/#
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/place
    /record/metadata/xMetaDiss/degree/grantor/universityOrInstitution/place/#
    /record/metadata/xMetaDiss/degree/level
    /record/metadata/xMetaDiss/degree/level/#
    /record/metadata/xMetaDiss/fileNumber
    /record/metadata/xMetaDiss/fileNumber/#
    /record/metadata/xMetaDiss/issued
    /record/metadata/xMetaDiss/issued/#
    /record/metadata/xMetaDiss/issued/@type
    /record/metadata/xMetaDiss/language
    /record/metadata/xMetaDiss/language/#
    /record/metadata/xMetaDiss/language/@type
    /record/metadata/xMetaDiss/modified
    /record/metadata/xMetaDiss/modified/#
    /record/metadata/xMetaDiss/modified/@type
    /record/metadata/xMetaDiss/publisher
    /record/metadata/xMetaDiss/publisher/@type
    /record/metadata/xMetaDiss/publisher/address
    /record/metadata/xMetaDiss/publisher/address/#
    /record/metadata/xMetaDiss/publisher/address/@Scheme
    /record/metadata/xMetaDiss/publisher/universityOrInstitution
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/department
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/department/name
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/department/name/#
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/name
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/name/#
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/place
    /record/metadata/xMetaDiss/publisher/universityOrInstitution/place/#
    /record/metadata/xMetaDiss/rights/#
    /record/metadata/xMetaDiss/rights/@kind
    /record/metadata/xMetaDiss/rights/@type
    /record/metadata/xMetaDiss/source
    /record/metadata/xMetaDiss/source/#
    /record/metadata/xMetaDiss/source/@type
    /record/metadata/xMetaDiss/transfer
    /record/metadata/xMetaDiss/transfer/#
    /record/metadata/xMetaDiss/transfer/@type
    []/record/header/setSpec
    []/record/header/setSpec/#
    []/record/metadata/xMetaDiss/abstract
    []/record/metadata/xMetaDiss/abstract/#
    []/record/metadata/xMetaDiss/abstract/@lang
    []/record/metadata/xMetaDiss/abstract/@type
    []/record/metadata/xMetaDiss/contributor
    []/record/metadata/xMetaDiss/contributor/@role
    []/record/metadata/xMetaDiss/contributor/@type
    []/record/metadata/xMetaDiss/contributor/person
    []/record/metadata/xMetaDiss/contributor/person/name
    []/record/metadata/xMetaDiss/contributor/person/name/@type
    []/record/metadata/xMetaDiss/contributor/person/name/personEnteredUnderGivenName
    []/record/metadata/xMetaDiss/contributor/person/name/personEnteredUnderGivenName/#
    []/record/metadata/xMetaDiss/creator
    []/record/metadata/xMetaDiss/creator/@type
    []/record/metadata/xMetaDiss/creator/person
    []/record/metadata/xMetaDiss/creator/person/name
    []/record/metadata/xMetaDiss/creator/person/name/@type
    []/record/metadata/xMetaDiss/creator/person/name/foreName
    []/record/metadata/xMetaDiss/creator/person/name/foreName/#
    []/record/metadata/xMetaDiss/creator/person/name/personEnteredUnderGivenName
    []/record/metadata/xMetaDiss/creator/person/name/personEnteredUnderGivenName/#
    []/record/metadata/xMetaDiss/creator/person/name/surName
    []/record/metadata/xMetaDiss/creator/person/name/surName/#
    []/record/metadata/xMetaDiss/fileProperties
    []/record/metadata/xMetaDiss/fileProperties/@fileID
    []/record/metadata/xMetaDiss/fileProperties/@fileName
    []/record/metadata/xMetaDiss/fileProperties/@fileSize
    []/record/metadata/xMetaDiss/identifier
    []/record/metadata/xMetaDiss/identifier/#
    []/record/metadata/xMetaDiss/identifier/@type
    []/record/metadata/xMetaDiss/isPartOf
    []/record/metadata/xMetaDiss/isPartOf/#
    []/record/metadata/xMetaDiss/isPartOf/@type
    []/record/metadata/xMetaDiss/rights
    []/record/metadata/xMetaDiss/subject
    []/record/metadata/xMetaDiss/subject/#
    []/record/metadata/xMetaDiss/subject/@type
    []/record/metadata/xMetaDiss/title
    []/record/metadata/xMetaDiss/title/#
    []/record/metadata/xMetaDiss/title/@lang
    []/record/metadata/xMetaDiss/title/@type
    []/record/metadata/xMetaDiss/type
    []/record/metadata/xMetaDiss/type/#
    []/record/metadata/xMetaDiss/type/@type

A dumb schema inferring
[visitor](https://asciinema.org/a/3x6u9yv2nsyqaf85evuu27tav). 1 single, 2
repeated.

    $ xmlray -visitor g -path /article test.xml 2> /dev/null
    {
      "/article": 1,
      "/article/@article-type": 1,
      "/article/@dtd-version": 1,
      "/article/@lang": 1,
      "/article/@mml": 1,
      "/article/@xlink": 1,
      "/article/@xsi": 1,
      "/article/front": 1,
      "/article/front/abstract": 1,
      "/article/front/article-meta": 1,
      "/article/front/article-meta/#": 1,
      "/article/front/article-meta/abstract": 2,
      "/article/front/article-meta/abstract/@abstract-type": 1,
      "/article/front/article-meta/abstract/@lang": 1,
      "/article/front/article-meta/abstract/disp-formula": 2,
      "/article/front/article-meta/abstract/disp-formula/tex-math": 2,
      "/article/front/article-meta/abstract/disp-formula/tex-math/#": 2,
      "/article/front/article-meta/abstract/disp-formula/tex-math/@id": 1,
      "/article/front/article-meta/abstract/disp-quote": 2,
      "/article/front/article-meta/abstract/fig": 1,
      "/article/front/article-meta/abstract/fig/@id": 1,
      "/article/front/article-meta/abstract/fig/@orientation": 1,
      "/article/front/article-meta/abstract/fig/@position": 1,
      "/article/front/article-meta/abstract/fig/caption": 1,
      "/article/front/article-meta/abstract/fig/caption/title": 1,
      "/article/front/article-meta/abstract/fig/graphic": 1,
      "/article/front/article-meta/abstract/fig/graphic/@href": 1,
      "/article/front/article-meta/abstract/fig/graphic/@orientation": 1,
      "/article/front/article-meta/abstract/fig/graphic/@position": 1,
      "/article/front/article-meta/abstract/fig/graphic/@type": 1,
      "/article/front/article-meta/abstract/label": 2,
      "/article/front/article-meta/abstract/label/#": 2,
      "/article/front/article-meta/abstract/list": 2,
      "/article/front/article-meta/abstract/list/@list-type": 1,
      "/article/front/article-meta/abstract/list/list-item": 2,
      "/article/front/article-meta/abstract/list/list-item/label": 2,
      "/article/front/article-meta/abstract/list/list-item/label/#": 2,
      "/article/front/article-meta/abstract/sec": 2,
      "/article/front/article-meta/abstract/sec/label": 2,
      "/article/front/article-meta/abstract/sec/label/#": 2,
      "/article/front/article-meta/abstract/statement": 1,
      "/article/front/article-meta/abstract/title": 2,
      "/article/front/article-meta/abstract/title/#": 2,
      "/article/front/article-meta/abstract/title/sc": 1,
      "/article/front/article-meta/abstract/title/sc/#": 1,
      "/article/front/article-meta/aff": 2,
      "/article/front/article-meta/aff/#": 2,
      "/article/front/article-meta/aff/@id": 1,
      "/article/front/article-meta/aff/country": 2,
      "/article/front/article-meta/aff/country/#": 2,
      "/article/front/article-meta/aff/email": 2,
      "/article/front/article-meta/aff/email/#": 2,
      "/article/front/article-meta/aff/email/@href": 1,
      "/article/front/article-meta/aff/email/@xlink": 1,
      "/article/front/article-meta/aff/institution": 2,
      "/article/front/article-meta/aff/institution/#": 2,
      "/article/front/article-meta/aff/institution/@type": 1,
      "/article/front/article-meta/aff/label": 2,
      "/article/front/article-meta/aff/label/#": 2,
      "/article/front/article-meta/aff/label/sup": 2,
      "/article/front/article-meta/aff/label/sup/#": 2,
      "/article/front/article-meta/aff/sc": 1,
      "/article/front/article-meta/aff/sc/#": 1,
      "/article/front/article-meta/aff/sup": 2,
      "/article/front/article-meta/aff/sup/#": 2,
      "/article/front/article-meta/article-categories": 1,
      "/article/front/article-meta/article-categories/series-text": 1,
      "/article/front/article-meta/article-categories/series-text/#": 1,
      "/article/front/article-meta/article-categories/series-title": 2,
      "/article/front/article-meta/article-categories/series-title/#": 2,
      "/article/front/article-meta/article-categories/series-title/sc": 2,
      "/article/front/article-meta/article-categories/series-title/sc/#": 2,
      "/article/front/article-meta/article-categories/subj-group": 2,
      "/article/front/article-meta/article-categories/subj-group/@subj-group-type": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/@subj-group-type": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/@subj-group-type": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/subj-group": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/subj-group/subject": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/subj-group/subject/#": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/subject": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subj-group/subject/#": 1,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject/#": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject/sc": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject/sc/#": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject/sup": 2,
      "/article/front/article-meta/article-categories/subj-group/subj-group/subject/sup/#": 2,
      "/article/front/article-meta/article-categories/subj-group/subject": 2,
      "/article/front/article-meta/article-categories/subj-group/subject/#": 2,
      "/article/front/article-meta/article-categories/subj-group/subject/@content-type": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/inline-graphic": 2,
      "/article/front/article-meta/article-categories/subj-group/subject/inline-graphic/@href": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/inline-graphic/@xlink": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/inline-graphic/alt-text": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/inline-graphic/alt-text/#": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/sc": 2,
      "/article/front/article-meta/article-categories/subj-group/subject/sc/#": 2,
      "/article/front/article-meta/article-categories/subj-group/subject/sup": 1,
      "/article/front/article-meta/article-categories/subj-group/subject/sup/#": 1,
      "/article/front/article-meta/article-id": 2,
      "/article/front/article-meta/article-id/#": 2,
      "/article/front/article-meta/article-id/@pub-id-type": 1,
      "/article/front/article-meta/article-id/@xlink": 1,
      ...
      "/article/sub-article/front/journal-meta": 2,
      "/article/sub-article/front/journal-meta/issn": 2,
      "/article/sub-article/front/journal-meta/issn/#": 2,
      "/article/sub-article/front/journal-meta/issn/@pub-type": 1,
      "/article/sub-article/front/journal-meta/journal-id": 2,
      "/article/sub-article/front/journal-meta/journal-id/#": 2,
      "/article/sub-article/front/journal-meta/journal-id/@journal-id-type": 1,
      "/article/sub-article/front/journal-meta/journal-title-group": 2,
      "/article/sub-article/front/journal-meta/journal-title-group/journal-title": 2,
      "/article/sub-article/front/journal-meta/journal-title-group/journal-title/#": 2,
      "/article/sub-article/front/journal-meta/publisher": 2,
      "/article/sub-article/front/journal-meta/publisher/publisher-name": 2,
      "/article/sub-article/front/journal-meta/publisher/publisher-name/#": 2
    }

What else?
----------

* Small number of values, like less than hundred?
* Is it a number, a string, a name?
* A [typical](https://golang.org/doc/codewalk/markov/) string?
* Typical length?
* Language?
* Frequency distribution?
* Element cooccurence?

How could some generated code look like?

Input
-----

    /record
    /record/header
    /record/header/datestamp
    /record/header/datestamp/#
    /record/header/identifier
    /record/header/identifier/#
    /record/metadata
    /record/metadata/xMetaDiss
    /record/metadata/xMetaDiss/@aiiso

Output
------

    type Record struct {
        xml.Name `xml:"record"`
        Header struct {
            xml.Name          `xml:"header"`
            Datestamp  string `xml:"datestamp"`
            Identifier string `xml:"identifier"`
        }
        Metadata struct {
            xml.Name `xml:"metadata"`
            XMetaDiss struct {
                xml.Name     `xml:"xMetaDiss"`
                Aiiso string `xml:"aiiso,attr"`
            }
        }
    }

The smallest unit would probably be:

    {{ .Name }} {{ if .IsList }} [] {{ end }} struct {
        xml.Name `xml:"{{ .Tag }}"`
        {{ range .Children }}
            // self
        {{ end }}
    }
