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

class ListTVC: UITableViewController {
    
    var operationTitle: String?
    var method: String?
    
    private var tableData: JSON = JSON.nullJSON
    
    @IBOutlet private weak var list: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.title = operationTitle
        
        fetchData()
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
            cell.textLabel?.text = self.tableData[indexPath.row]["fechahora"].string!
            cell.detailTextLabel?.text = self.tableData[indexPath.row]["nombremedicamento"].string! + " | Quantitat: " + self.tableData[indexPath.row]["cantidad"].string!
            break;
        case "Salidas":
            cell.textLabel?.text = self.tableData[indexPath.row]["nombrefarmacia"].string! + " | " + self.tableData[indexPath.row]["fechahora"].string!
            cell.detailTextLabel?.text = self.tableData[indexPath.row]["nombremedicamento"].string! + " | Quantitat: " + self.tableData[indexPath.row]["cantidad"].string!
            break;
        case "Medicamentos":
            cell.textLabel?.text = self.tableData[indexPath.row]["nombre"].string!
            let estoc = self.tableData[indexPath.row]["enalmacen"].string!
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
            }
        }
    }
    
    // MARK: - Table animation
    func animateTable() {
        tableView.reloadData()
        
        let cells = tableView.visibleCells()
        let tableHeight = tableView.bounds.size.height
        
        // Move all cells to the bottom of the screen
        for c in cells {
            let cell = c as UITableViewCell
            cell.transform = CGAffineTransformMakeTranslation(0, tableHeight)
        }
        
        var index = 0
        for c in cells {
            let cell = c as UITableViewCell
            UIView.animateWithDuration(1.5, delay: 0.05 * Double(index), usingSpringWithDamping: 0.8, initialSpringVelocity: 0, options: nil, animations: {
                cell.transform = CGAffineTransformMakeTranslation(0, 0)
            }, completion: nil)
            
            index++
        }
        
    }
    
    // MARK: - API Client
    
    private func fetchData() {
        var url: String = GlobalConstants.apiEndpoint + method!
        
        Alamofire.request(.GET, url, parameters: nil)
            .responseJSON { (req, res, json, error) in
                if(error != nil) {
                    NSLog("Error: \(error)\nfor url: \(url)")
                }
                else {
                    NSLog("Success: \(url)")
                    self.tableData = JSON(json!)
                    self.animateTable()
                }
        }
    }
}
