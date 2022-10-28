//使用此包前，必须已经启动了config包了
package notice

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/wujiyu98/ginframe/model"
)

func Message(m model.Message) {
	var buf bytes.Buffer
	html := `<div>
    <table style="width: 100%;">
      <tr>
        <th style="text-align: left; width: 150px;padding: 5px;">Email:</th>
        <td>{{.Email}}</td>
      </tr>
      <tr>
        <th style="text-align: left; width: 150px;padding: 5px;">Name:</th>
        <td>{{.Name}}</td>
      </tr>

      <tr>
        <th style="text-align: left; width: 150px;padding: 5px;">Phone:</th>
        <td>{{.MobilePhone}}</td>
      </tr>
      <tr>
        <th style="text-align: left; width: 150px;padding: 5px;">Country:</th>
        <td>{{.Country}}</td>
      </tr>
      <tr>
        <th style="text-align: left; width: 150px;padding: 5px;">Company:</th>
        <td>{{.Company}}</td>
      </tr>
    </table>
    <div style="padding: 10px;">
      <p>{{.Comment}}</p>
    </div>
  </div>`
	t, _ := template.New("message").Parse(html)
	t.Execute(&buf, m)
	subject := fmt.Sprintf("Website Message-%s", m.Name)

	email.send(m.Email, subject, buf.String())

}

func Enquiry(m model.Enquiry) {
	var buf bytes.Buffer
	html := `  <div>
    <h2 style="font-size: 1.5em;padding: 10px 0">Request Quote</h2>
    <table width="100%" border="1" cellspacing="0" cellpadding="10"
      style="border: 1px solid #6c757d;border-collapse: collapse;">
      <thead>
        <tr>
          <th style="text-align: left">Part Number</th>
          <th style="text-align: left">Manufacturers</th>
          <th style="text-align: left">Request Qty</th>
          <th style="text-align: left">Target Price</th>
          <th style="text-align: left">Description</th>
        </tr>
      </thead>
      <tbody>
        {{range $item := .Products}}
        <tr>
          <td>{{$item.Title}}</td>
          <td>{{$item.Manufacturer}}</td>
          <td>{{$item.Qty}}</td>
          <td>{{$item.Price}}</td>
          <td>{{$item.Summary}}</td>
        </tr>
        {{end}}
      </tbody>
    </table>
    <br>
    
    <p><b>Country：</b>{{.Country}}</p>
    <p><b>Company：</b>{{.Company}}</p>
    <p><b>Name：</b>{{.Name}}</p>
    <p><b>Mobile Phone：</b>{{.MobilePhone}}</p>
    <p><b>Email：</b>{{.Email}}</p>
    <p>{{.Comment}}</p>
  </div>`
	t, _ := template.New("enquiry").Parse(html)
	t.Execute(&buf, m)
	subject := fmt.Sprintf("Website Enquiry-%s", m.Name)

	email.send(m.Email, subject, buf.String())

}
