//
//  OperationTVC.swift
//  Client API
//
//  Created by Andrés Pizá on 20/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit

class OperationTVC: UITableViewController {
    
    // API operations
    let operations = ["Categorias", "Entradas", "Salidas", "Medicamentos", "Noticias", "Farmacias"]

    //TODO  let operations = {"Categorias": "/categorias", "Entradas", "Salidas", "Medicamentos", "Noticias", "Farmacias"}
    
    override func viewDidLoad() {
        super.viewDidLoad()
    }
    
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    // MARK: - Table view data source
    
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return operations.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("cell", forIndexPath: indexPath) as UITableViewCell
        
        cell.textLabel?.text = operations[indexPath.row]
        
        return cell
    }
    
    override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
        if indexPath.section == 0 {
            performSegueWithIdentifier("showList", sender: indexPath)
        }
    }
    
    override func prepareForSegue(segue: UIStoryboardSegue, sender: AnyObject?) {
        if let tvc = segue.destinationViewController as? ListTVC {
            if let index = sender as? NSIndexPath {
                tvc.operationTitle = operations[index.row]
            }
        }
    }
}
