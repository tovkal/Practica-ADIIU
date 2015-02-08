//
//  ListTVC.swift
//  Client API
//
//  Created by Andrés Pizá on 20/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit
import SwiftyJSON
import Alamofire

class ListTVC: UITableViewController, UITableViewDataSource {
    
    var operationTitle: String?
    var method: String?
    
    private var tableData: JSON = JSON.nullJSON
    
    @IBOutlet private weak var list: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.title = operationTitle
        
        fetchData()
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    // MARK: - Table view data source
    
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return tableData.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("listCell", forIndexPath: indexPath) as UITableViewCell
        
        // PEr tenir diferents custom cells, emprar identificadors diferents
        
        switch (self.operationTitle!) {
        case "Categorias":
            cell.textLabel?.text = self.tableData[indexPath.row]["nombre"].string!
            cell.detailTextLabel?.text = self.tableData[indexPath.row]["texto"].string!
            break;
        case "Entradas":
            cell.textLabel?.text = "Medicament: " + self.tableData[indexPath.row]["idmedicamento"].string! + " | Quantitat: " + self.tableData[indexPath.row]["cantidad"].string!
            cell.detailTextLabel?.text = self.tableData[indexPath.row]["fechahora"].string!
            break;
        case "Salidas":
            cell.textLabel?.text = "Medicament: " + self.tableData[indexPath.row]["idmedicamento"].string! + " | Quantitat: " + self.tableData[indexPath.row]["cantidad"].string!
            cell.detailTextLabel?.text = "Farmacia: " + self.tableData[indexPath.row]["idfarmacia"].string! + " | Data: " + self.tableData[indexPath.row]["fechahora"].string!
            break;
        case "Medicamentos":
            cell.textLabel?.text = self.tableData[indexPath.row]["nombre"].string!
            let estoc = self.tableData[indexPath.row]["enalmacen"].number!
            cell.detailTextLabel?.text = "Estoc: \(estoc)"
            break;
        case "Noticias":
            cell.textLabel?.text = self.tableData[indexPath.row]["texto"].string!
            cell.detailTextLabel?.text = "Inici: " + self.tableData[indexPath.row]["inicio"].string! + " | Fi: " + self.tableData[indexPath.row]["fin"].string!
            break;
        case "Farmacias":
            cell.textLabel?.text = "Nick: " + self.tableData[indexPath.row]["nik"].string! + " | Pass: " + self.tableData[indexPath.row]["pass"].string!
            cell.detailTextLabel?.text = "Nivell: " + self.tableData[indexPath.row]["nivel"].string!
            break;
        default:
            break;
        }
        
        return cell
    }
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if indexPath.section == 0 {
            performSegueWithIdentifier("showDetail", sender: indexPath)
        }
    }
    
    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
        if let dvc = segue.destinationViewController as? DetailVC {
            if let index = sender as? NSIndexPath {
                dvc.operation = self.operationTitle
                dvc.data = self.tableData[index.row]
                // Pasar dades
                // Pasar tipus dades
            }
        }
    }
    
    // MARK: - API Client
    
    private func fetchData() {
        var url: String = GlobalConstants.api + method!
        
        Alamofire.request(.GET, url, parameters: nil)
            .responseJSON { (req, res, json, error) in
                if(error != nil) {
                    NSLog("Error: \(error)\nfor url: \(url)")
                }
                else {
                    NSLog("Success: \(url)")
                    self.tableData = JSON(json!)
                    self.list.reloadData()
                }
        }
    }
}
