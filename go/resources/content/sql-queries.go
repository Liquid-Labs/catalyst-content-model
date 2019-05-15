package content

// Content
// create - queries
const CreateContentQuery = `INSERT INTO content (id, extern_path, slug, type, title, summary, version_cookie) VALUES(?,?,?,?,?,?,?)`
const CreateContentTypeTextQuery = `INSERT INTO content_type_text (id, format, text) VALUES(?,?,?)`
// retrieve - partials
const CommonContentFields = `e.pub_id, e.last_updated, c.title, c.summary, c.namespace, c.slug, c.type `
const CommonCententTypeTextFields = `ctt.format, ctt.text, c.extern_path, c.last_sync, c.version_cookie `
const CommonContentContribFields = `p.pub_id, p.display_name, pc.role, pc.summary_credit_order `
const CommonContentFrom = `FROM content c JOIN contributors pc ON c.id=pc.content JOIN persons p ON cp.id=p.id `
const CommonContentTypeTextFields = CommonContentFields + `, ctt.format, ctt.text `
const CommonContentTypeTextFrom = `JOIN ctt content_type_text ON c.id=ctt.id `
const CommonContentTypeTextGet string = `SELECT ` + CommonContentTypeTextFields + CommonContentFrom + CommonContentTypeTextFrom
//  - queries - partials
const GetContentTypeTextQuery string = CommonContentTypeTextGet + `WHERE e.pub_id=? `
const GetContentTypeTextByNSSlugQuery string = CommonContentTypeTextGet + ` WHERE c.namespace=? AND c.slug=? `
const GetContentTypeTextByIDQuery string = CommonContentTypeTextGet + ` WHERE c.id=? `
// update - queries
const UpdateContentTypeTextSansTextQuery = `UPDATE content c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, c.title=?, c.summary=?, c.extern_path=?, c.namespace=?, c.slug=?, ctt.format=? WHERE e.pub_id=?`
const UpdateContentTypeTextWithTextQuery = `UPDATE content c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, ctt.last_sync=0, c.title=?, c.summary=?, c.extern_path=?, c.namespace=?, c.slug=?, ctt.format=?, ctt.text=? WHERE e.pub_id=?`
const UpdateContentTypeTextOnlyTextQuery = `UPDATE content c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, ctt.last_sync=0, ctt.text=? WHERE e.pub_id=?`

// Contributors
// update by refresh only
const ContributorsDeleteQuery = `DELETE * FROM contributors WHERE content=?`
const ContributorInsertQuery = `INSERT INTO contributors (id, content, role, summary_credit_order) SELECT persons.id, content.id, ?, ? FROM persons p JOIN content c ON p.pub_id=? AND c.pub_id=?`
const ContributorInsertWithContentIDQuery = `INSERT INTO contributors (id, content, role, summary_credit_order) SELECT persons.id, ?, ?, ? FROM persons p WHERE p.pub_id=?`
