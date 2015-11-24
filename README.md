README
======

xmlray is a little x-ray things for xml.

![](http://etc.usf.edu/clipart/22900/22922/ray_22922_sm.gif)

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
