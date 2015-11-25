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
